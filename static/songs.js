var songApp = angular.module("songApp", []);

songApp.controller("songCtrl", function($scope) {
    $scope.artist = "Mathew";
    $scope.songs = ['song 1', 'song 2'];
});