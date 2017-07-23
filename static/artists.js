var artistApp = angular.module("artistApp", []);


function getArtists($scope) {
    $.get("get_artists", (artists, error) => {
        $scope.$apply(() => 
            $scope.artists = artists);
    });
}

artistApp.controller("artistCtrl", function($scope) {
    getArtists($scope);
});