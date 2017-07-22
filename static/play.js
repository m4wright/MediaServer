var songApp = angular.module("playSong", []);

function song_info(song_path) {
    let path_to_song = song_path.split("/");
    if (path_to_song.length == 0) {
        return "";
    }
    let song_name = path_to_song[path_to_song.length - 1];
    let artist = path_to_song[2];
    let extensionIndex = song_name.lastIndexOf(".");
    if (extensionIndex === -1) {
        return song_name;
    }
    return {
        name: song_name.substr(0, extensionIndex),
        artist: artist,
        path: song_path
    };
}

function getSongs($scope) {
    $.get("get_songs", (songs, error) => {
        $scope.$apply(() => 
        $scope.songs = songs.map(song_path => song_info(song_path)));
    });
}

songApp.controller("playCtrl", function($scope) {
    getSongs($scope);
});