package markdown

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type MarkdownOutput struct {
	Html    template.HTML
	Headers []*MarkdownHeader
}

type MarkdownHeader struct {
	Level    int               `json:"level"`
	Title    string            `json:"title"`
	Slug     string            `json:"slug"`
	Children []*MarkdownHeader `json:"children"`
}

func RenderMarkdownFile(path string) (*MarkdownOutput, error) {
	// read in markdown file
	md, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return RenderMarkdown(md)
}

func RenderMarkdown(md []byte) (*MarkdownOutput, error) {
	// parse as markdown
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	// renderer := html.NewRenderer(opts)
	renderer := &customRenderer{
		Renderer: *html.NewRenderer(opts),
	}

	rendered := markdown.Render(doc, renderer)

	if rendered == nil {
		return nil, fmt.Errorf("there was an issue rendering the markdown")
	}

	// get headers
	headers := generateHeaders(md)

	return &MarkdownOutput{
		Html:    template.HTML(rendered),
		Headers: headers,
	}, nil
}

type customRenderer struct {
	html.Renderer
}

func (cr *customRenderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	if node, ok := node.(*ast.Heading); ok && entering {
		// Generate slug from the text content of the header
		slug := createSlug(string(node.Children[0].AsLeaf().Literal))
		node.HeadingID = slug
	}

	// Fall back to default rendering
	return cr.Renderer.RenderNode(w, node, entering)
}

func createSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)
	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove any non-alphanumeric (excluding hyphen)
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, slug)
	return slug
}

func generateHeaders(md []byte) []*MarkdownHeader {
	var headers []*MarkdownHeader
	stack := []*MarkdownHeader{}

	scanner := bufio.NewScanner(bytes.NewReader(md))
	for scanner.Scan() {
		line := scanner.Text()
		level := strings.LastIndex(line, "#") + 1
		if level > 0 {
			title := strings.TrimSpace(line[level:])
			header := &MarkdownHeader{Level: level, Title: title, Slug: createSlug(title)}
			for len(stack) > 0 && stack[len(stack)-1].Level >= level {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				headers = append(headers, header)
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, header)
			}
			stack = append(stack, header)
		}
	}
	return headers
}
