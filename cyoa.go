package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Story : Holds a bunch of chapters
type Story map[string]Chapter

// Chapter : holds a number of paragraphs and options.
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option : Text to display, Arcs to point to.
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var storyTpl, indexTpl *template.Template

type handler struct {
	s Story
	t *template.Template
}

func init() {
	indexTpl = template.Must(template.ParseFiles("web/index.html"))
	storyTpl = template.Must(template.ParseFiles("web/chapter.html"))
}

// NewHandler : Turns a story into an HTTP handler
func NewHandler(s Story, t *template.Template) http.Handler {
	if t == nil {
		t = storyTpl
	}
	return handler{s, t}
}

// ServeHTTP : Method for handlers, takes in a writer and a request and serves a web page
func (h handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Get chapter from URL
	path := strings.TrimSpace(req.URL.Path)
	if path != "" || path != "/" {
		path = path[1:] // Strips the '/' prefix

		if chapter, ok := h.s[path]; ok {
			err := storyTpl.ExecuteTemplate(res, "chapter.html", chapter)
			if err != nil {
				log.Printf("%v", err)
				http.Error(res, "Something went wrong :C\n", http.StatusInternalServerError)
			}
			return
		}
		http.Error(res, "Chapter not found :C\n", http.StatusNotFound)

	}
	path = "/"

	err := indexTpl.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		log.Printf("%v", err)
		http.Error(res, "Something went wrong :C\n", http.StatusInternalServerError)
	}
}

// JSONStory : Takes a reader, decodes from JSON into Story
func JSONStory(reader io.Reader) (Story, error) {
	decoder := json.NewDecoder(reader)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
