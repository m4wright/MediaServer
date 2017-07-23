var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("http://192.168.0.134:8080/get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => 
            $scope.songs = songs);
    });
}

function getKeys(obj) {
    return Object.keys(obj);
}

songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
    $scope.keys = getKeys;
});