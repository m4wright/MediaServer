var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("http://192.168.0.134:8080/get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => 
            $scope.songs = songs);
            // $scope.song_names = Object.keys(songs);
            // console.log("keys: " + JSON.stringify($scope.song_names));
    });
}


songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
    $scope.song_names = ['Never Thought That This Would Happen'];
});