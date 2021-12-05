package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"story"
)

func main() {
	port := flag.Int("port", 3000, "the port to start web application")
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story %s \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	storyfirst, err := story.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := story.NewHandler(storyfirst)
	fmt.Printf("Starting server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
