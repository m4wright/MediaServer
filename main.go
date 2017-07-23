package main

import (
	"net/http"
)

var path_to_music = "/home/mathew/go/MediaServer/music"


func handleMusic(w http.ResponseWriter, r *http.Request) {
	// when playing a song, redirect to the actual path of the song
	new_path := path_to_music + r.URL.Path[len("/music/"):]
	http.ServeFile(w, r, new_path)
}



func main() {
	generate_songs(path_to_music)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/music/", handleMusic)
	http.HandleFunc("/upload_file", upload)
	http.HandleFunc("/songs", choose_song_html)
	http.HandleFunc("/upload", upload_html)
	http.HandleFunc("/get_songs", get_songs_html)
	http.ListenAndServe(":8080", nil)
}
