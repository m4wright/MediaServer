var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("../get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => {
            $scope.songs = songs;
            $scope.song_names = Object.keys(songs);
            $scope.current_song_path = $scope.songs[$scope.song_names[0]];
        });
    });
}




songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
    $scope.play_song = function(song) {
        console.log("new song name: " + song);
        $scope.current_song_path = $scope.songs[song];
        console.log("new song path: " + $scope.current_song_path);
    };
});