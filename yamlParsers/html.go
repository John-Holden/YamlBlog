package yamlParser

import (
	"fmt"
	"log"
	"strings"
)

func GetPostListBodyHtml(post_dir string) string {
	var htmlString strings.Builder
	htmlString.WriteString((markdownLink("🏚️ JohnHolden", "/")))
	htmlString.WriteString(ListPosts(post_dir, ""))
	return GetBody(htmlString.String())
}

func GetPostHtml(filename string) string {
	fmt.Printf("[i] Rendering post %s \n", filename)
	data := loadYaml(filename)
	head := GetPostHeadHTML(data)
	body := GetPostBodyHTML(data)

	body = GetBody(body)
	return GetPage(head, body)
}

func GetPostHeadHTML(data map[string]interface{}) string {
	favicon := "../static/favicon.ico"
	title := getInterfaceStrAttr(data["title"])

	if title == "" {
		log.Fatalf("A a title is required, found none")
	}

	description := getInterfaceStrAttr(data["description"])
	css_paths := getInterfaceSliceAttr(data["css_paths"])

	if css_paths == nil {
		css_paths = []string{"../css/default.css"}
	}

	json_paths := getInterfaceSliceAttr(data["json_paths"])

	if json_paths == nil {
		json_paths = []string{}
	}

	return GetHead(
		title,
		description,
		favicon,
		css_paths,
		json_paths,
	)
}

// Renders a given post element by element and returns HTML-formatted string
func GetPostBodyHTML(data map[string]interface{}) string {
	var htmlString strings.Builder
	// Nav
	htmlString.WriteString(markdownLink("🔙", "/"))

	// Write post title
	title_md := []byte("# " + getInterfaceStrAttr(data["title"]))
	htmlString.WriteString(string(mdToHTML(title_md)))

	// Write post elements
	body_html, err := GetBodyHtml(data)

	if err != nil {
		log.Fatalf("Error parsing head block: %s\n", err)
	}

	for _, element := range body_html {
		htmlString.WriteString(element)
	}

	return htmlString.String()
}