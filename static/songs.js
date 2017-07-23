var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("http://192.168.0.134:8080/get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => 
            $scope.songs = songs);
            getSongNames($scope);
    });
}


function getSongNames($scope) {
    let song_names = [];
    for (let song in $scope.songs) {
        song_names.push($scope.songs[song]);
    }
    console.log(song_names);
    $scope.$apply(() => {
        $scope.song_names = song_names;
    });
}

songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
});