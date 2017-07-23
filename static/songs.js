var songApp = angular.module("songApp", []);

var songs;
var song_names;

function getSongs($scope) {
    $.post("http://192.168.0.134:8080/get_songs", {artist: document.title}, (song_list, error) => {
        $scope.$apply(() => {
            $scope.songs = song_list;
            songs = song_list;
            $scope.song_names = Object.keys(songs);
            song_names = $scope.song_names;
            $scope.current_song_path = "http://192.168.0.134:8080" + $scope.songs[$scope.song_names[0]].substr(1);
        });
    });
}

function nextSong(song_path) {
    song_path = decodeURI(song_path);
    console.log('path: ' + song_path);
    console.log("other: " + songs[song_names[0]]);
    for (let i = 0; i < song_names.length; i++) {
        if (song_path.indexOf(songs[song_names[i]].substr(1)) >= 0) {
            console.log("matched. Next: " + songs[song_names[(i + 1) % song_names.length]]);
            return songs[song_names[(i + 1) % song_names.length]];
        }
    }
    return songs[song_names[0]];
}


songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
    $scope.play_song = function(song) {
        $scope.current_song_path = encodeURI("http://192.168.0.134:8080" + $scope.songs[song].substr(1));
        return $scope.current_song_path;
    };
    document.getElementById("audio").addEventListener('ended', function() {
        console.log("ended");
        this.src = $scope.play_song(nextSong(this.src));
        console.log("source: " + this.src);
        this.play();
    });
});