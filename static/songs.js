var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("http://192.168.0.134:8080/get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => 
            $scope.songs = songs);
            $scope.song_names = getKeys(songs);
            console.log("keys: " + $scope.song_names);
    });
}

function getKeys(obj) {
    return Object.keys(obj);
}

songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
});