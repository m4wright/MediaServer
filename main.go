package main

import (
	"net/http"
)




func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/music/", http.StripPrefix("/music/", http.FileServer(http.Dir("music"))))
	http.HandleFunc("/upload_file", upload)
	http.HandleFunc("/songs", choose_song_html)
	http.HandleFunc("/upload", upload_html)
	http.HandleFunc("/get_songs", get_songs_html)
	http.ListenAndServe(":8080", nil)
}
