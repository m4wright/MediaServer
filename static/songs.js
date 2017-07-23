var songApp = angular.module("songApp", []);

songApp.controller("songCtrl", function($scope) {
        $scope.songs = ['song 1', 'song 2'];
});

var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => 
            $scope.songs = songs);
    });
}

artistApp.controller("artistCtrl", function($scope) {
    getSongs($scope);
});