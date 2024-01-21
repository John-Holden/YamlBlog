# YamBlog - A Blog Application with YAML to HTML Conversion

YamBlog is a simple blog application written in Go that allows you to create and serve blog pages from YAML files. Posts are stored in the `/pages` dir. Go converts each YAML element  to HTML for rendering. 

This README will guide you through setting up and using the application.

## Features - WIP
- Serve blog pages from YAML files
- Convert Markdown content to HTML for rendering
- Support for serving static assets like CSS files
- CORS headers and content-type handling
- Integration with Google Cloud Functions Framework

## Installation

1. Install the project and dependencies
```Bash
git clone https://github.com/John-Holden/yamBlog.git
cd yamBlog
# Install deps
go mod tidy
# Build the app
go build
```
2. Set up your blog content:
    - Create a directory named pages where you'll store your YAML blog content files.
    - Add YAML files in the pages directory. Each YAML file represents a blog page. You can structure your YAML files as shown in the provided examples.
3. Set up static assets (CSS):
    - Create a directory named static for your static assets, including CSS files.
    - Place your CSS files in the static directory.

## Usage

Once you have set up the YamBlog application, you can use it to serve your blog content.

1. Start the local web server:
    ```Bash
    cd yamBlog
    go run main.go
    ```
    This will start the local web server on port 8080. You can customize the port by setting the PORT environment variable.
2. Access your blog:
    - To access the list of blog posts, open your web browser and go to:
    http://localhost:8080
    - To access individual blog posts, use the following URL pattern: 
        http://localhost:8080/post/{post-name}.
3. Customize your blog:
    - You can customize your blog by modifying the content of the pages directory (YAML files) and the CSS files in the static directory.

## YAML Content Format
Each YAML file in the pages directory represents a blog page and should adhere to the following schema:

```Yaml
title: MyTitle # Renders as a <h1> html & also embedded in <head>
description: A brief introduction into how to parse yaml etc # embedded in <head>
body:
  # Currently don't do anything with date !TODO!
  - date: "2023-03-01"
  # Currently don't do anything with author !TODO!
  - author: JohnHolden
  # Markdown input text rendered to HTML using github.com/gomarkdown/markdown pkg
  - text: | 
      ## This is a second sub-title
      
      Lorem Lorem dolor sit amet, consectetur 
      adipiscing elit. Nullam 
      Pellentesque habitant morbi tristique senectus et netus et malesuada
      fames ac turpis egestas.

      ---
      
      Vivamus eget scelerisque nulla. Fusce ac leo vel enim luctus rhoncus.
  - text: |
      - Lorem
      - Lorem
      - dolor
  # Include code snippets following the below pattern:
  - code:
      lang: python
      input: |
        # This is python
        print('hello world')
        for i in range(10):
          print(i)
  - code:
      lang: go
      input: |
        // This is GO
        fmt.Println("


```

## Future Work - this really is a quick and dirty first approximation
- Render Date/Author
- Proper validation of input yaml file
- tests + actions
- Deployment to cloud + test hosting on serverless