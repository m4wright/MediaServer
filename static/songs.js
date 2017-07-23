var songApp = angular.module("songApp", []);


function getSongs($scope) {
    $.post("http://192.168.0.134:8080/get_songs", {artist: document.title}, (songs, error) => {
        $scope.$apply(() => {
            $scope.songs = songs;
            $scope.song_names = Object.keys(songs);
            $scope.current_song_path = "http://192.168.0.134:8080" + $scope.songs[$scope.song_names[0]].substr(1);
        });
    });
}

function play_song(song) {
    $scope.$apply(() => 
        $scope.current_song_path = $scope.songs[song]);
}

// songApp.filter("trustUrl", ['$sce', function($sce) {
//     return function(recordingUrl) {
//         return $sce.trustAsResourceUrl(recordingUrl);
//     }
// }])

songApp.controller("songCtrl", function($scope) {
    getSongs($scope);
    $scope.play_song = function(song) {
        $scope.current_song_path = $scope.songs[song];
        console.log($scope.current_song_path);
    };
});