package server

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// parseTemplates handles the logic behind parsing templates and attaching functions.
func (s *server) parseTemplates(w http.ResponseWriter, files ...string) (tpl *template.Template, err error) {
	// Automatically adds app layout
	files = append(files, filepath.Join("layouts", "app"))
	for i, v := range files {
		files[i] = filepath.Join("client", "templates", v) + ".tpl"
	}
	// Automatically adds components folder
  comps, err := ioutil.ReadDir("." + filepath.Join("client", "templates", "components"))
  if err != nil {
    s.logErr("failed to read components", err)
    return nil, err
  }
	for _, v := range comps {
    files = append(files, v.Name())
  }

	tpl, err = template.New("").Funcs(template.FuncMap{
		"echo": func(input string) string {
			return input
		},
		"isMarkdown": func(data interface{}) bool {
			switch data.(type) {
			case markdown:
				return true
			default:
				return false
			}
		},
	}).ParseFiles(files...)
	if err != nil {
		s.logErr("Error parsing template file", err)
		return nil, err
	}
	return tpl, nil
}

func (s *server) parseData(data interface{}) interface{} {
	switch data.(type) {
	case nil:
		return nil
	default:
		return data
	}
}

type markdown struct {
	Body     *bytes.Buffer
	MetaData map[string]interface{}
}

// parseMarkdown will take in a markdown file location and return html
func (s *server) parseMarkdown(file string) (markdown, error) {
	var (
		buf bytes.Buffer
		err error
	)
	file = file + ".md"

	src, err := ioutil.ReadFile(file)
	if err != nil {
    s.logErr("failed to read .md file", err)
    return markdown{}, err
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	ctx := parser.NewContext()
	err = md.Convert([]byte(src), &buf, parser.WithContext(ctx))
	if err != nil {
    s.logErr("markdown failed to parse", err)
    return markdown{}, err
	}

	// TODO validate MetaData so that I can ensure that markdown files are valid posts.
	return markdown{
		Body:     &buf,
		MetaData: meta.Get(ctx),
	}, nil
}
