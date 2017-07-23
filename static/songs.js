var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("../get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => 
            $scope.songs = songs);
    });
}

songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
});