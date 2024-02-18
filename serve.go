package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/John-Holden/YamlBlog/Parsers"
	"github.com/gorilla/mux"
)

// Global constants

type BlogConf struct {
	content,
	static,
	port,
	css string
}

func SetBlogConf(
	content_path string,
	static_path string,
	port int,
	css_path string) *BlogConf {
	return &BlogConf{
		content: content_path,
		static:  static_path,
		port:    fmt.Sprint(port),
		css:     css_path,
	}
}

func DefaultConf() *BlogConf {
	return SetBlogConf(
		"pages",
		"static",
		8080,
		"css",
	)
}

// Set CSS headers
func CSSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		next(w, r)
	}
}

// Set CORS headers
func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// If it's a preflight OPTIONS request, return here
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}

// Serve static CSS files
func ServeCSS(w http.ResponseWriter, r *http.Request) {
	cssFilePath := r.URL.Path[1:]
	cssContent, err := ioutil.ReadFile(cssFilePath)
	if err != nil {
		fmt.Println("Error reading CSS file:", err)
		return
	}
	fmt.Fprintf(w, string(cssContent))
}

// Renders a single post
func RenderPost(w http.ResponseWriter, r *http.Request, contentDir string) {
	fmt.Println("[i] Rendering Post: " + r.URL.Path)
	html_doc := ""
	links := Parsers.GetPostPaths(contentDir)
	for link, filename := range links {
		if "/"+contentDir+"/"+link == r.URL.Path {
			html_doc = Parsers.GetPostHtml(contentDir + "/" + filename)
			break
		}
	}

	if html_doc == "" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}

	fmt.Fprintf(w, html_doc)
}

// Renders list of posts
func RenderPostList(w http.ResponseWriter, r *http.Request, conf BlogConf) {
	fmt.Println("[i] Rendering Post List...")

	HeadTitle := "Blog List"
	HeadDescription := "List of blog posts"

	var js_paths []string
	var css_paths = []string{conf.css + "/default.css"}

	head := Parsers.GetHead(
		HeadTitle,
		HeadDescription,
		conf.static+"flavicon.ico",
		css_paths,
		js_paths)

	body := Parsers.GetPostListBodyHtml(conf.content)
	fmt.Fprintf(w, Parsers.GetPage(head, body))

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func SetRoutes(w http.ResponseWriter, r *http.Request, conf BlogConf) {
	fmt.Println("[i] Setting WebServer routes...")
	router := mux.NewRouter()

	setRenderList := func(w http.ResponseWriter, r *http.Request) {
		RenderPostList(w, r, conf)
	}

	setRender := func(w http.ResponseWriter, r *http.Request) {
		RenderPost(w, r, conf.content)
	}

	router.HandleFunc("/", CorsMiddleware(setRenderList))
	router.HandleFunc("/"+conf.content+"/{post}", CorsMiddleware(setRender))
	router.HandleFunc("/"+conf.css+"/{css}", CSSMiddleware(ServeCSS))
	router.PathPrefix(conf.static).Handler(
		http.StripPrefix(
			conf.static,
			http.FileServer(
				http.Dir("static"),
			),
		),
	)
	router.ServeHTTP(w, r)
}
