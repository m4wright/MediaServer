var songApp = angular.module("playSong", []);


function getSongs($scope) {
    $.get("get_artists", (songs, error) => {
        $scope.$apply(() => 
            $scope.songs = songs);
    });
}

songApp.controller("playCtrl", function($scope) {
    getSongs($scope);
});