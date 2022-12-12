package main

import (
	"log"
	"os"

	go_files "github.com/duexcoast/go_files"
)

func main() {
	posts, err := go_files.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
