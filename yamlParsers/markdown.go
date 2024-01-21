package yamlParser

import (
	"bytes"

	"github.com/alecthomas/chroma/quick"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func syntaxHighlight(
	input,
	lang,
	formatter,
	style string) (string, error) {
	// Convert language to syntax highlighted html
	var buf bytes.Buffer

	err := quick.Highlight(&buf, input, lang, formatter, style)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
