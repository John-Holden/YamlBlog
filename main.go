package YamlBlog

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	yamlP "github.com/John-Holden/YamlBlog/parsers"
	"github.com/gorilla/mux"
)

// Global constants
const (
	cssDir         = "css/default"
	contentDir     = "pages"
	staticDir      = "/static/"
	defaultFavIcon = "static/favicon.ico"
)

// Set CSS headers
func CSSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		next(w, r)
	}
}

// Set CORS headers
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
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

// Renders list of posts
func RenderPostList(w http.ResponseWriter, r *http.Request) {
	HeadTitle := "Blog List"
	HeadDescription := "List of blog posts"

	var js_paths []string
	css_paths := []string{"css/default.css"}

	head := yamlP.GetHead(
		HeadTitle,
		HeadDescription,
		defaultFavIcon,
		css_paths,
		js_paths)

	body := yamlP.GetPostListBodyHtml(contentDir)
	fmt.Fprintf(w, yamlP.GetPage(head, body))

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

// Renders a single post
func renderPost(w http.ResponseWriter, r *http.Request) {
	html_doc := ""
	links := yamlP.GetPostPaths(contentDir)
	for link, filename := range links {
		if "/post/"+link == r.URL.Path {
			html_doc = yamlP.GetPostHtml(contentDir + "/" + filename)
			break
		}
	}

	if html_doc == "" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}

	fmt.Fprintf(w, html_doc)
}

func setRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[i] Starting WebServer...")
	router := mux.NewRouter()
	router.HandleFunc("/", corsMiddleware(RenderPostList))
	router.HandleFunc("/post/{post}", corsMiddleware(renderPost))
	router.HandleFunc("/css/{css}", CSSMiddleware(ServeCSS))
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("static"))))
	router.ServeHTTP(w, r)
}

func start(
	port int,
	content_dir string,
	css_dir string,
	static_dir string,
	flavicon string)

func main() {
	fmt.Println("[i] Starting Local WebServer...")
	funcframework.RegisterHTTPFunction("/", setRoutes)
	// Use the PORT environment variable, or default to 8080
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
