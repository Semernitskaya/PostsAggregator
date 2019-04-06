package main

import (
	"fmt"
	"log"
	"net/http"
)

var interestsHolder = InterestsHolder{make(map[string]bool)}

func handlerAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", interestsHolder.String())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	interestsHolder.Load()
	http.HandleFunc("/all", handlerAll)
	for key := range interestsHolder.Cache {
		http.HandleFunc("/"+key, handler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
