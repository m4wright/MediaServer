package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"encoding/json"
	"os/exec"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func execute(args []string) string {
	cmd := exec.Command(args[0], args[1:]...)

	output, error := cmd.CombinedOutput()

	if error != nil {
		fmt.Println("Error executing command: " + error.Error())
		panic(error.Error())
	}

	return string(output)
}

func upload_html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/upload.html")
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
	fmt.Fprintf(w, get_songs("/home/mathew/Documents/Go/upload"))
}

func upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		panic(err.Error())
	}
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	err = r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	artist := r.Form.Get("artist")
	if artist == "" {
		artist = "no_artist"
	}
	if !exists("./music" + artist) {
		err = os.Mkdir("./music/"+artist, 0755)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Fprintf(w, "%v", handler.Header)

	if err != nil {
		panic(err.Error())
	}
	f, err := os.OpenFile("./music/"+artist+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
	}
	io.Copy(f, file)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/music/", http.StripPrefix("/music/", http.FileServer(http.Dir("music"))))
	http.HandleFunc("/upload_file", upload)
	http.HandleFunc("/songs", choose_song_html)
	http.HandleFunc("/upload", upload_html)
	http.HandleFunc("/get_songs", get_songs_html)
	http.ListenAndServe(":8080", nil)
}
