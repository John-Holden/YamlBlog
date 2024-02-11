//Generates HTML templates

package Parsers

import (
	"fmt"
	"log"
	"strings"
	"text/template"
)

// MetaInfo represents the information needed for a meta tag
type MetaInfo struct {
	Name    string
	Content string
}

type HeadData struct {
	Title           string
	MetaDescription string
	FavIcon         string
	CSSPaths        []string
	JSPaths         []string
}

// BodyData holds the dynamic content for the body
type BodyData struct {
	Content string
}

// PageData combines head and body data
type PageData struct {
	Head string
	Body string
}

// Get tile HTML tag
func GetTitleTag(content string) string {
	type Title struct {
		Title string
	}

	title := Title{
		Title: content,
	}

	tpl := template.Must(template.New("").Parse("<title>{{.Title}}</title>"))
	var sb strings.Builder

	if err := tpl.Execute(&sb, title); err != nil {
		panic(err)
	}

	return sb.String()
}

func GetMetaTag(content string) string {
	// Data for the meta tag
	meta := MetaInfo{
		Name:    "description",
		Content: content,
	}

	// Define a template string for a meta tag
	tmpl := `<meta name="{{.Name}}" content="{{.Content}}">`

	// Parse the template string
	t, err := template.New("metaTag").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Execute the template with the meta tag data
	var sb strings.Builder
	err = t.Execute(&sb, meta)
	if err != nil {
		panic(err)
	}

	return sb.String()
}

func GetHead(
	title string,
	description string,
	favicon string,
	css_paths []string,
	js_paths []string) string {

	data := HeadData{
		Title:           title,
		MetaDescription: description,
		FavIcon:         favicon,
		CSSPaths:        css_paths,
		JSPaths:         js_paths,
	}

	// Define the template
	tmpl := `
    <head>
        <meta charset="UTF-8">
        <title>{{.Title}}</title>
        <meta name="description" content="{{.MetaDescription}}">
		<link rel="icon" href="{{.FavIcon}}" type="image/x-icon">
        {{range .CSSPaths}}<link rel="stylesheet" href="{{.}}">{{end}}
        {{range .JSPaths}}<script src="{{.}}"></script>{{end}}
    </head>
    `

	// Parse and execute the template
	t, err := template.New("head").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	var sb strings.Builder
	err = t.Execute(&sb, data)
	if err != nil {
		log.Fatal(err)
	}

	return sb.String()
}

// Takes html-formatted header and content and
// creates a body tag
func GetBody(content string) string {
	data := BodyData{
		Content: content,
	}

	tmpl := `
    <body>
		<div class="container">
        	{{.Content}}
		</div>
    </body>
    `

	t, err := template.New("body").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	var sb strings.Builder
	err = t.Execute(&sb, data)
	if err != nil {
		log.Fatal(err)
	}

	return sb.String()
}

// Wraps around html-formatted header/body strings
// Returns html-formatted page document
func GetPage(head string, body string) string {
	data := PageData{
		Head: head,
		Body: body,
	}
	fmt.Println("[i] Rendering Page...")

	tmpl := `
    <!DOCTYPE html>
    <html>
        {{.Head}}
        {{.Body}}
    </html>
    `
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	var sb strings.Builder
	err = t.Execute(&sb, data)
	if err != nil {
		log.Fatal(err)
	}

	return sb.String()
}
