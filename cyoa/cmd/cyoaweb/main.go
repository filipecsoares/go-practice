package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/filipecsoares/go-practice/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// Open the JSON file and parse the story in it.
	f, err := os.Open(*filename)
	if err != nil {
		exit(err)
	}
	defer f.Close()
	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		exit(err)
	}
	fmt.Printf("%+v\n", story)
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
