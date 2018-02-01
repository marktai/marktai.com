marktai.controller("BombCtl", ["$scope", "$rootScope", "$route", "$sce", "$window", function($scope, $rootScope, $route, $sce, $window) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "bombmanual";
	var url = "https://www.marktai.com/download/me/Bomb-Defusal-Manual_1_stripped.pdf";
	$scope.google_url = $sce.trustAsResourceUrl("https://docs.google.com/viewer?url=" + url);
	$scope.url = $sce.trustAsResourceUrl(url);
	var pages = ["Simple Wires", "Colored Button", "Symbol Keypads", "Simon Says", "Word Display/Buttons", "Number Display/Buttons (Memory)", "Morse Code", "Complicated Wires", "Wire Sequences", "Mazes", "Passwords", "Venting Gas", "Capacitor", "Knobs", "Appendix"]
	$scope.text = "Index: \n" + pages.map(function(s){return '    '+s}).join('\n');
}])
