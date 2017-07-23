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
    for (let i = 0; i < song_names.length; i++) {
        if (songs[song_names[i]] === song_path) {
            console.log("matched");
            return songs[song_names[(i + 1) % song_names.length]];
        }
    }
    return songs[song_names[0]];
}


songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
    $scope.play_song = function(song) {
        $scope.current_song_path = encodeURI("http://192.168.0.134:8080" + $scope.songs[song].substr(1));
    };
    document.getElementById("audio").addEventListener('ended', function() {
        console.log("hi");
        $scope.play_song(song_names[1]);
        let audio = this;
        setTimeout(() => audio.play(), 1000);
        //this.play();
    });
});