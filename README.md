go programs to have:
main.go 		--> 		has the main function. Handles requests
upload.go 		--> 		handles the upload APIs. Used to upload songs,
							albums, etc

download.go		--> 		used to download a song off of youtube

get_songs.go 	-->			creates a map of artist: [songs]
							only called on startup. 
							Maybe creates a JSON of the map as well?
							if a new song is added, handles adding it to
							the map



