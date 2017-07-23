var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("http://192.168.0.134:8080/get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => 
            $scope.songs = songs);
    });
}


function getKeys(obj) {
    let keys = [];
    console.log("getKeys called");
    for (let key of obj) {
        if (obj.hasOwnProperty(key)) {
            keys.push(key);
        }
    }
    console.log(keys);
    return keys;
}

songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
    $scope.$apply(() => 
        $scope.getKeys = getKeys);
});