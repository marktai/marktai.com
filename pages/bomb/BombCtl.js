marktai.controller("BombCtl", ["$scope", "$rootScope", "$route", "$sce", "$window", function($scope, $rootScope, $route, $sce, $window) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "bombmanual";
	var url = "https://www.marktai.com/download/me/Bomb-Defusal-Manual_1.pdf";
	$scope.google_url = $sce.trustAsResourceUrl("https://docs.google.com/viewer?url=" + url);
	$scope.url = $sce.trustAsResourceUrl(url);
}])
