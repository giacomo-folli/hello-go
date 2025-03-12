package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/giacomo-folli/hello-go/17_blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		first = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		second = `Title: Post 1
Description: Description 1
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(first)},
		"hello-world2.md": {Data: []byte(second)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
