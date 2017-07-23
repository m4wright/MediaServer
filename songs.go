package main

import (
	"net/http"
	"html/template"
	"strings"
	"fmt"
)

func handle_songs(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split_path := strings.Split(path, "/")	
	if len(split_path) < 3 {
		panic("Missing artist")
	}
	artist := split_path[2]

	songsTemplate, err := template.New("songs.html").Delims("<<<", ">>>").ParseFiles("./templates/songs.html")
	if err != nil {
		panic(err.Error())
	}

	err = songsTemplate.Execute(w, artist)
	if err != nil {
		panic(err.Error())
	}
}