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
	if len(split_path) < 2 {
		panic("Missing artist")
	}
	artist := split_path[1]
	fmt.Println("artist: " + artist)

	t, err := template.ParseFiles("./templates/songs.html")
	if err != nil {
		panic(err.Error())
	}

	t.Execute(w, artist)
}