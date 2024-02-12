package main

import (
	"github.com/John-Holden/YamlBlog/WebServer"
)

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
