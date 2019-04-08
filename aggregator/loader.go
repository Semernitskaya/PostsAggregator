package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type PostsLoader interface {
	LoadPostsByTheme(theme string, count int) []string
}

type MockPostsLoader struct {
	posts map[string][]string
}

func (loader MockPostsLoader) LoadPostsByTheme(theme string, count int) []string {
	theme = "#" + theme
	posts := loader.posts[theme]
	if len(posts) < count {
		return posts
	}
	return posts[0:count]
}

func New() *MockPostsLoader {
	return &MockPostsLoader{map[string][]string{
		"#java":   {"j1", "j2"},
		"#golang": {"g1", "g2", "g3"},
		"#devops": {"", ""},
		"#else":   {"e1", "e2"},
	}}
}

type NewsApiLoader struct{}

func (loader NewsApiLoader) LoadPostsByTheme(theme string, count int) []string {
	resp, err := http.Get("https://newsapi.org/v2/everything?" +
		"q=" + theme +
		"&pageSize=" + strconv.Itoa(count) +
		"&from=2019-04-08" +
		"&sortBy=popularity" +
		"&apiKey=" + os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return []string{string(bodyBytes)}
}
