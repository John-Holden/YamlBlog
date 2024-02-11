# YamBlog - A Blog Application with YAML to HTML Conversion

YamBlog is a simple blog application written in Go that allows you to create and serve HTML pages from markdown contained in YAML files. HTML pages are stored in the `/pages` dir by default. Each YAML element is then converted into HTML. 

This README will guide you through setting up and using the application.

## Features - WIP
- Serve super simple blog pages from YAML files
- Convert Markdown content to HTML for rendering
- Support for serving static assets like CSS files
- CORS headers and content-type handling

## Module Usage
- TODO

## Running Locally

- Install the project and dependencies:
```Bash
git clone https://github.com/John-Holden/YamlBlog.
cd YamlBlog
go mod tidy
```
- Run the webserver using the default config:
```
go run main.go
```
- List all HTML pages: `http://localhost:8080/`
- Navigate to a page, e.g.: `http://localhost:8080/pages/basic`
- Update `/pages` with new HTML pages according to the schema
- The default paths can be updated by calling the `WebServer.SetBlogConf` method in main.go:

```GO
func main() {
	content_dir := "posts"
	static_dir := "files"
	css_dir := "css"
	port := 80
	WebServer.Start(
		WebServer.SetBlogConf(
			content_dir,
			static_dir,
			port,
			css_dir,
		),
	)
}
```

## YAML Content Format
Each YAML file in the pages directory represents a blog page and should adhere to the following schema:

```Yaml
# Renders 'title' as a <h1> html & also embedded in <head>
title: MyTitle 
# embedded in <head>
description: A brief introduction into how to parse yaml etc
body: # Rendered as <body>
  # Currently don't do anything with date/author !TODO!
  - date: "2023-03-01"
  - author: JoeBlogs

  # Markdown input text rendered to HTML using github.com/gomarkdown/markdown pkg
  - text: | 
      ## This is a second sub-title
      
      Lorem Lorem dolor sit amet, consectetur 
      adipiscing elit. Nullam 
      Pellentesque habitant morbi tristique senectus et netus et malesuada
      fames ac turpis egestas.      
      Vivamus eget scelerisque nulla. Fusce ac leo vel enim luctus rhoncus.
  # Include code snippets following the below pattern:
  - code:
      lang: python
      input: |
        # This is python
        print('hello world')
        for i in range(10):
          print(i)

```

## Future Work - this really is a quick and dirty first approximation
- Render Date/Author
- Proper validation of input yaml file
- tests + actions
- Deployment to cloud + test hosting on serverless