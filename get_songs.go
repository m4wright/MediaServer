package main

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
	"os/exec"
)



func execute(args []string) string {
	cmd := exec.Command(args[0], args[1:]...)

	output, error := cmd.CombinedOutput()

	if error != nil {
		fmt.Println("Error executing command: " + error.Error())
		panic(error.Error())
	}

	return string(output)
}



func choose_song_html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/play.html")
}

func get_songs(base_path string) string {
	songsString := execute([]string{"/usr/bin/find", base_path, "-name", "*.mp3"})
	songsString = strings.Trim(songsString, "\r\n\t")
	songsLocation := strings.Split(songsString, "\n")

	length_of_base := len(base_path)
	for i := 0; i < len(songsLocation); i++ {
		songsLocation[i] = songsLocation[i][length_of_base:]
	}

	songsEncoding, err := json.Marshal(songsLocation)
	if err != nil {
		panic(err.Error())
	}
	if songsEncoding == nil {
		panic("Error encoding json")
	}

	return string(songsEncoding)
}

func get_songs_html(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, get_songs("/home/mathew/Documents/Go/upload"))
	fmt.Fprintf(w, get_songs("/home/mathew/go/MediaServer")
}

