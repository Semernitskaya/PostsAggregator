package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var interestsHolder = InterestsHolder{make(map[string]bool)}

var loaders = []PostsLoader{New(), NewsApiLoader{}}

func handlerAll(w http.ResponseWriter, r *http.Request) {
	var posts []string
	for interest := range interestsHolder.Cache {
		posts = append(posts, loadPosts(interest)...)
	}
	fmt.Fprintf(w, "Posts %s! %s", interestsHolder.String(), strings.Join(posts, " | "))
}

func handler(w http.ResponseWriter, r *http.Request) {
	interest := r.URL.Path[1:]
	posts := loadPosts(interest)
	fmt.Fprintf(w, "Posts %s! %s", interest, strings.Join(posts, " | "))
}

func loadPosts(interest string) []string {
	var posts []string
	for _, loader := range loaders {
		posts = append(posts, loader.LoadPostsByTheme(interest, 10)...)
	}
	return posts
}

func main() {
	interestsHolder.Load()
	http.HandleFunc("/all", handlerAll)
	for interest := range interestsHolder.Cache {
		http.HandleFunc("/"+interest, handler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
