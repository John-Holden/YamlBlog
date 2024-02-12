package Parsers

import (
	"fmt"
	"reflect"
	"regexp"
)

type ParserFunc func(interface{}) (string, error)

var HeadParsers = map[string]ParserFunc{
	"title":       Title,
	"description": Description,
	"tags":        Tags,
}

var BodyParsers = map[string]ParserFunc{
	"text":   Text,
	"author": Author,
	"date":   Date,
	"code":   Code,
}

// Do nothing
func Type(input interface{}) (string, error) {
	fmt.Println("[i] Parsing Type")
	return "", nil
}

func Title(input interface{}) (string, error) {
	fmt.Println("[i] Parsing Title")
	v := getInterfaceStrAttr(input)
	if v == "" {
		return "", fmt.Errorf("Unsupported type %v", reflect.TypeOf(v))
	}

	title_tag := GetTitleTag(v)
	return title_tag, nil
}

func Text(input interface{}) (string, error) {
	fmt.Println("[i] Parsing Text")
	md := []byte(getInterfaceStrAttr(input))
	return string(mdToHTML(md)), nil
}

func Description(input interface{}) (string, error) {
	fmt.Println("[i] Parsing Meta")
	v := getInterfaceStrAttr(input)
	if v == "" {
		return "", fmt.Errorf("Unsupported type %v", reflect.TypeOf(v))
	}

	meta_html := GetMetaTag(v)
	return meta_html, nil
}

func Date(input interface{}) (string, error) {
	fmt.Println("[i] Parsing Date")
	return "", nil
}

func Tags(input interface{}) (string, error) {
	fmt.Println("[i] Parsing Tags")
	return "", nil
}

func Author(input interface{}) (string, error) {
	fmt.Println("[i] Parsing Author")
	return "", nil
}

func Code(input interface{}) (string, error) {
	fmt.Println("[i] Parsing Code...")
	code := ""
	lexer := ""
	format := "html"
	style := "dracula"

	switch input.(type) {
	case string:
		code = getInterfaceStrAttr(input)
	case map[string]interface{}:
		code = getInterfaceStrAttr(input.(map[string]interface{})["input"])
		lexer = getInterfaceStrAttr(input.(map[string]interface{})["lang"])
	}

	syntaxHiHtml, err := syntaxHighlight(
		code,
		lexer,
		format,
		style)

	if err != nil {
		return "", err
	}

	// Remove redundant stand-alone HTML
	re := regexp.MustCompile(`body\s*{[^}]*}`)
	cleanedString := re.ReplaceAllString(syntaxHiHtml, "")
	re = regexp.MustCompile(`|<html>|</html>`)
	cleanedString = re.ReplaceAllString(cleanedString, "")

	re = regexp.MustCompile(`body`)
	cleanedString = re.ReplaceAllString(cleanedString, "div")

	return cleanedString, nil
}
