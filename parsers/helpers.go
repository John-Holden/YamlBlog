package Parsers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v3"
)

func markdownHTMLconverter(input string) string {
	return string(blackfriday.Run([]byte(input)))
}

// Takes list of strings and builds markdown-valued
// string before converting into a HTML-valued string
func markdownList(list []string) string {
	var markdownList strings.Builder
	for _, item := range list {
		markdownList.WriteString(fmt.Sprintf("- %s\n", item))
	}
	return markdownHTMLconverter(markdownList.String())
}

func markdownLink(text string, url string) string {
	return markdownHTMLconverter(fmt.Sprintf("[%s](%s)", text, url))
}

// Finds and returns list of all posts
// & displays as html-formatted string
func ListPosts(dir string, host string) string {
	fmt.Println("[i] Listing Posts...")
	filenames := FindDirFiles(dir, "yaml")
	post_links := GetPostLinks(dir, filenames, host)
	return markdownList(post_links)
}

func DirFileNames(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var file_names []string
	for _, file := range files {
		if file.Type().IsDir() {
			continue
		}
		file_names = append(file_names, file.Name())
	}
	return file_names
}

func FilterFileType(file_names []string, file_type string) []string {
	var yaml_files []string
	for _, file := range file_names {
		if !strings.HasSuffix(file, file_type) {
			continue
		}
		yaml_files = append(yaml_files, file)
	}
	return yaml_files
}

func FindDirFiles(dir string, filetype string) []string {
	return FilterFileType(DirFileNames(dir), filetype)
}

// Loads yaml from input filename string
// returns generic map of all yaml data
func loadYaml(filename string) map[string]interface{} {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %s\n", err)
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	return data
}

func loadHead(data map[string]interface{}) (map[string]string, error) {
	var head map[string]string
	head = make(map[string]string)

	head_input, ok := data["head"]
	if !ok {
		return nil, errors.New("Could not load head section.")
	}

	switch head_ := head_input.(type) {
	case map[string]interface{}:
		for k, v := range head_ {
			head_str, err := HeadParsers[k](v)
			if err != nil {
				return nil, errors.New("Failed to parse head element: " + k)
			}
			head[k] = head_str
		}
	}
	return head, nil
}

func GetBodyHtml(data map[string]interface{}) ([]string, error) {
	body := []string{}

	body_input, ok := data["body"]
	if !ok {
		return nil, errors.New("Could not load body section.")
	}

	switch body_ := body_input.(type) {
	case []interface{}:
		for _, body_element := range body_ {
			switch element := body_element.(type) {
			case map[string]interface{}:
				for element_k, element_v := range element {
					element_str, err := BodyParsers[element_k](element_v)
					if err != nil {
						return nil, errors.New("Could not load body element: " + element_k)
					}
					body = append(body, element_str)
				}
			}
		}
	default:
		return nil, errors.New("Could not parse Body.")
	}
	return body, nil

}

// Get map of links and filenames
func GetPostPaths(dir string) map[string]string {
	pathMaps := make(map[string]string)
	file_ext := ".yaml"
	filenames := FindDirFiles(dir, file_ext)
	for _, filename := range filenames {
		pathMaps[strings.Replace(filename, file_ext, "", -1)] = filename
	}
	return pathMaps
}

// Get map of html links for all posts
func GetPostLinks(dir string, filenames []string, host string) []string {
	var postLinks []string
	file_ext := ".yaml"
	for _, filename := range filenames {
		data := loadYaml(dir + "/" + filename)
		path := strings.Replace(filename, file_ext, "", -1)
		title, err := getMapStrAttr("title", data)
		if err != nil {
			return nil
		}
		postLinks = append(postLinks, markdownLink(title, host+"/"+dir+"/"+path))
	}
	return postLinks
}

// getMapStrAttr extracts a string attribute from various types of input
func getMapStrAttr(attrName string, data map[string]interface{}) (string, error) {
	v, ok := data[attrName]
	if !ok {
		return "", errors.New("type key not found in data")
	}

	switch v := v.(type) {
	case string:
		// Handle the case where the input is a string
		return v, nil
	default:
		return "", fmt.Errorf("Unsupported type %v", reflect.TypeOf(v))
	}
}

// getMapStrAttr checks if the input is a string and returns, else returns an nil.
func getInterfaceStrAttr(input interface{}) string {
	switch v := input.(type) {
	case string:
		return v
	default:
		return ""
	}
}

func getInterfaceSliceAttr(input interface{}) []string {
	switch v := input.(type) {
	case []string:
		return v
	default:
		return nil
	}
}
