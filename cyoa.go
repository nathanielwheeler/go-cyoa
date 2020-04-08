package cyoa

import (
	"log"
	"net/http"
	"html/template"
	"encoding/json"
	"io"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(""))
}

// NewHandler : Casts story into an HTTP handler
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

// TODO export into its own file
// ServeHTTP : 
func (h handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	http.HandleFunc("/", index)
	http.Handle("/web/", http.StripPrefix("/web", http.FileServer(http.Dir("web"))))
	http.ListenAndServe(":8080", nil)
}



func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

// JSONStory : Takes a reader, decodes from JSON into Story
func JSONStory(reader io.Reader) (Story, error){
	decoder := json.NewDecoder(reader)
	var story Story
	if err:= decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

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
