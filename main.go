package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func Start(config *BlogConf) {
	fmt.Println("[i] Starting Local WebServer...")
	if config == nil {
		config = DefaultConf()
	}

	setRoutesWithConfig := func(w http.ResponseWriter, r *http.Request) {
		SetRoutes(w, r, *config)
	}

	funcframework.RegisterHTTPFunction("/", setRoutesWithConfig)

	if err := funcframework.Start(config.port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

func main() {
	content_dir := "pages"
	static_dir := "static"
	css_dir := "css"
	port := 80

	Start(
		SetBlogConf(
			content_dir,
			static_dir,
			port,
			css_dir,
		),
	)
}
