package go_files

import "testing/fstest"

type Post struct {
}

func NewPostsFromFS(filesystem fstest.MapFS) []Post {
	return []Post{{}, {}}
}
