package main

import (
	"strings"
)

type InterestsHolder struct {
	Cache map[string]bool
}

func (h *InterestsHolder) String() string {
	var str strings.Builder
	for key := range h.Cache {
		str.WriteString("#" + key + " ")
	}
	return str.String()
}

func (h *InterestsHolder) Load() {
	h.Cache["golang"] = true
	h.Cache["java"] = true
	h.Cache["devops"] = true
}

func (h *InterestsHolder) AddInterest(interest string) {
	//TODO: save data to DB
	//TODO: multithreading execution
	h.Cache[interest] = true
}
