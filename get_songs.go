package main

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
	"os/exec"
)


var songs map[string](map[string]string) 	
/*
	songs is of type: 
		{artist: {
			song_name: song_path
		}
	}
*/

var artists []string
// a slice of the artists (essentially the keys of songs)

var artists_string string

func execute(args []string) string {
	cmd := exec.Command(args[0], args[1:]...)

	output, error := cmd.CombinedOutput()

	if error != nil {
		fmt.Println("Error executing command: " + error.Error())
		panic(error.Error())
	}

	return string(output)
}

	
func get_song_and_artist(path string) (string, string) {
	/*
		returns songname, artist
		assumes the format is:
			/path/to/music/Artist/songname.mp3
	*/

	path_list := strings.Split(path, "/")
	if len(path_list) < 2 {
		panic("Invalid path to song")
	}
	song_with_extension := path_list[len(path_list) - 1]
	artist := path_list[len(path_list) - 2]

	indexOfExtension := strings.LastIndexAny(song_with_extension, ".")
	if indexOfExtension < 0 {
		panic("Invalid path to song")
	}
	song_name := song_with_extension[:indexOfExtension]

	return artist, song_name
}


func generate_song_list(base_path string) {
	// creates the songs map
	songs = make(map[string](map[string]string))

	songsString := execute([]string{"/usr/bin/find", base_path, "-name", "*.mp3"})
	songsString = strings.Trim(songsString, "\r\n\t")
	songsLocation := strings.Split(songsString, "\n")

	length_of_base := len(base_path)

	for i := 0; i < len(songsLocation); i++ {
		songsLocation[i] = "./music" + songsLocation[i][length_of_base:]
		artist, song_name := get_song_and_artist(songsLocation[i])
		if songs[artist] == nil {
			songs[artist] = make(map[string]string)
		}
		songs[artist][song_name] = songsLocation[i]
	}

	artists = make([]string, len(songs))

	i := 0
	for artist, _ := range(songs) {
		artists[i] = artist
		i++
	}

	string_of_artists, err := json.Marshal(artists)
	if err != nil {
		panic(err.Error())
	}

	artists_string = string(string_of_artists)
}


func choose_artist_html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/artists.html")
}

func get_artists_request(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, artists_string)
}