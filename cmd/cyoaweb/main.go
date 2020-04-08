package main

import (
	"log"
	"os"
	"flag"

	"github.com/nathanielwheeler/cyoa"
)

func main() {
	filename := flag.String("file", "story.json", "JSON file with CYOA story")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("File not found")
	}

	story, err := cyoa.JSONStory(file)
	if err != nil {
		log.Fatalln("Error decoding story file.")
	}
}