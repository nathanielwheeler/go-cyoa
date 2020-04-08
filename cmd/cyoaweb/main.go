package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"flag"

	"github.com/nathanielwheeler/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "CYOA starts on this port.")
	filename := flag.String("file", "story.json", "JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using story in %s\n", *filename)

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("File not found")
	}

	story, err := cyoa.JSONStory(file)
	if err != nil {
		log.Fatalln("Error decoding story file.")
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Serving on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}