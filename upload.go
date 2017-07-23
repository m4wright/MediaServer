package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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



func upload_html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/upload.html")
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
	if !exists("./music/" + artist) {
		err = os.Mkdir("./music/" + artist, 0755)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Fprintf(w, "%v", handler.Header)

	path := "./music/" + artist + "/" + handler.Filename

	indexOfExtension := strings.LastIndexAny(song_with_extension, ".")
	if indexOfExtension < 0 {
		panic("Invalid filename")
	}

	song_name := handler.Filename[:indexOfExtension]

	if songs[artist] == nil {
		songs[artist] = make(map[string]string)
	}
	songs[artist][song_name] = path

	if err != nil {
		panic(err.Error())
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
	}
	err = io.Copy(f, file)
	if err != nil {
		panic(err.Error())
	}
}

