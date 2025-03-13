package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	blogposts "github.com/giacomo-folli/hello-go/17_blogposts"
	blogrenderer "github.com/giacomo-folli/hello-go/18_blogrenderer"
)

const POSTS_DIR = "./posts"
const STATIC_DIR = "./static"

func StartServer() {
	http.Handle("/", http.FileServer(http.Dir(STATIC_DIR)))
	http.ListenAndServe(":8080", nil)
}

func LoadPosts() (posts []blogposts.Post) {
	posts, err := blogposts.NewPostsFromFS(os.DirFS(POSTS_DIR))
	if err != nil {
		log.Fatalf("failed to load posts: %v\n", err)
	}
	return
}

func ExportPosts(posts []blogposts.Post) {
	if err := os.Mkdir(STATIC_DIR, 0777); os.IsExist(err) {
		fmt.Println("\"static\" directory already present")
	} else if err != nil {
		log.Fatalf("failed to create \"static\" directory")
	} else {
		fmt.Println("Created \"static\" directory")
	}

	for _, p := range posts {
		postRenderer, err := blogrenderer.NewPostRender()
		if err != nil {
			log.Fatalf("failed to create post renderer: %v\n", err)
		}

		post := blogrenderer.Post{
			Title:       p.Title,
			Description: p.Description,
			Body:        p.Body,
			Tags:        p.Tags,
		}

		fileName := fmt.Sprintf("%s.html", post.SanitisedTitle())
		fmt.Println("Creating file with name:", fileName)

		f, err := os.Create(strings.Join([]string{STATIC_DIR, fileName}, "/"))
		if err != nil {
			log.Fatalf("failed to create file for post %s: %v\n", post.Title, err)
		}
		defer f.Close()

		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, post); err != nil {
			log.Fatalf("failed to render post %s: %v\n", post.Title, err)
		}

		if _, err := io.Copy(f, &buf); err != nil {
			log.Fatalf("failed to write post %s to file: %v\n", post.Title, err)
		}
	}
}

func main() {
	posts := LoadPosts()

	ExportPosts(posts)

	StartServer()
}
