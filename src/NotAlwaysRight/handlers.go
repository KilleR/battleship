package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"regexp"
)

func apiInit(router *mux.Router) {
	//router.HandleFunc("/", getAllDeviceDetailsHandler)
	router.HandleFunc("/{storyid}", getStoryHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The world is but a broken toy")
	fmt.Fprintln(w, "It's pleasures hollow, false it's joy")
	fmt.Fprintln(w, "I'm a config engine")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var pathParts = regexp.MustCompile("(/[^/]+)").FindAllString(r.URL.Path, -1)
	for _,v := range pathParts {
		fmt.Fprintln(w, v)
	}
}

func getStoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storyId := vars["storyid"]

	var stories []story
	var err error
	stories, err = parsePage("http://notalwaysright.com/-/"+storyId)
}