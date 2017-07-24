var songApp = angular.module("songApp", []);

var songs;
var song_names;

var url = "http://192.168.0.134:8080";

function getSongs($scope) {
    $.post("http://192.168.0.134:8080/get_songs", {artist: document.title}, (song_list, error) => {
        $scope.$apply(() => {
            $scope.songs = song_list;
            songs = song_list;
            $scope.song_names = Object.keys(songs);
            song_names = $scope.song_names;
            $scope.current_song_path = url + $scope.songs[$scope.song_names[0]].substr(1);
            $scope.current_song = $scope.song_names[0];
        });
    });
}

function nextSong(song_path) {
    song_path = decodeURI(song_path);
    for (let i = 0; i < song_names.length; i++) {
        if (song_path.indexOf(songs[song_names[i]].substr(1)) >= 0) {
            return song_names[(i + 1) % song_names.length];
        }
    }
    return song_names[0];
}


songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
    $scope.artist = document.title;
    $scope.play_song = function(song) {
        $scope.current_song = song;
        $scope.current_song_path = encodeURI(url + $scope.songs[song].substr(1));
        return $scope.current_song_path;
    };
    document.getElementById("audio").addEventListener('ended', function() {
        this.src = $scope.play_song(nextSong(this.src));
        this.play();
    });
});