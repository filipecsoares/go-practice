package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/filipecsoares/go-practice/cyoa"
)

func main() {
	// Create flags for our optional variables
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// Open the JSON file and parse the story in it.
	f, err := os.Open(*filename)
	if err != nil {
		exit(err)
	}
	defer f.Close()
	story, err := cyoa.JsonStory(f)
	if err != nil {
		exit(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
