package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var interestsHolder = InterestsHolder{make(map[string]bool)}

var twitterLoader = New()

func handlerAll(w http.ResponseWriter, r *http.Request) {
	var tweets []string
	for key := range interestsHolder.Cache {
		hashTag := "#" + key
		tweets = append(tweets, twitterLoader.LoadTweetsByHashTag(hashTag, 10)...)
	}
	fmt.Fprintf(w, "Tweets %s! %s", interestsHolder.String(), strings.Join(tweets, " | "))
}

func handler(w http.ResponseWriter, r *http.Request) {
	hashTag := "#" + r.URL.Path[1:]
	tweets := twitterLoader.LoadTweetsByHashTag(hashTag, 10)
	fmt.Fprintf(w, "Tweets %s! %s", hashTag, strings.Join(tweets, " | "))
}

func main() {
	interestsHolder.Load()
	http.HandleFunc("/all", handlerAll)
	for interest := range interestsHolder.Cache {
		http.HandleFunc("/"+interest, handler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
