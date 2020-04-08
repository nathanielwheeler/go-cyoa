package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
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

var tpl *template.Template

// TODO switch to templates
var handlerTemplate = `
<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css"
		integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
	<link rel="stylesheet" href="style.css">
	<title>¡CYOA!</title>
</head>

<body>
	<header>
		<!-- Chapter title goes here -->
		<h1 class="text-center">{{.Title}}</h1>
	</header>
	<main class="container">
		<div class="row">
			<!-- Paragraphs go here -->
			{{range .Paragraphs}}
			<p class="col-12">
				{{.}}
			</p>
			{{end}}
		</div>
		<div class="row">
			<!-- Options go here -->
			{{range .Options}}
			<div class="col">
				<button class="btn btn-outline-danger btn-lg btn-block"
					onclick="window.location.href = '/{{.Arc}}';">{{.Text}}</button>
			</div>
			{{end}}
		</div>
	</main>
</body>

</html>
`

type handler struct {
	s Story
}

func init() {
	tpl = template.Must(template.New("").Parse(handlerTemplate))
}

// NewHandler : Turns a story into an HTTP handler
func NewHandler(s Story) http.Handler {
	fmt.Println("Handling story")
	return handler{s}
}

// ServeHTTP : Method for handlers, takes in a writer and a request and serves a web page
func (h handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Executing story...")
	err := tpl.Execute(res, h.s["intro"])
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
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
