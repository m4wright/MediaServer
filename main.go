package main

import (
	"net/http"
	"fmt"
)

var path_to_music = "/home/mathew/go/MediaServer/music"


func handleMusic(w http.ResponseWriter, r *http.Request) {
	// when playing a song, redirect to the actual path of the song
	fmt.Println("original path: " + r.URL.Path)
	new_path := path_to_music + r.URL.Path[len("/music/"):]
	fmt.Println("new path: " + new_path)
	http.ServeFile(w, r, new_path)
}


func main() {
	generate_song_list(path_to_music)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/music/", handleMusic)
	http.HandleFunc("/upload_file", upload)
	http.HandleFunc("/artists", choose_artist_html)
	http.HandleFunc("/upload", upload_html)
	http.HandleFunc("/songs/", handle_songs)
	http.HandleFunc("/get_songs", get_songs)
	http.HandleFunc("/get_artists", get_artists_request)
	http.ListenAndServe(":8080", nil)
}
