package go_files

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}
	var posts []Post

	for _, f := range dir {
		post, err := getPost(filesystem, f)
		if err != nil {
			return nil, err // needs clarification: should we totally fail if one file fails, or just ignore?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := filesystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)

}

func newPost(postFile fs.File) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil

}
