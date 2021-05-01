package main

import (
	"bandersnatch/cyoa"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Story")
	file := flag.String("file", "gopher.json", "The JSON file with CYIA story")
	flag.Parse()
	fmt.Printf("Using the file %s\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		fmt.Println("File not found")
		os.Exit(1)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", story)

}
