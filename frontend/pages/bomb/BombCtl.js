marktai.controller("BombCtl", ["$scope", "$rootScope", "$route", "$sce", "$window", function($scope, $rootScope, $route, $sce, $window) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "bombmanual";
	var url = "https://www.marktai.com/download/54689/Bomb-Defusal-Manual_1.pdf";
	$scope.google_url = $sce.trustAsResourceUrl("https://docs.google.com/viewer?url=" + url);
	$scope.url = $sce.trustAsResourceUrl(url);
	var pages = ["Simple Wires - 5", "Colored Button - 6", "Symbol Keypads - 7", "Simon Says - 8", "Word Display - 9/10", "Number Display - 11", "Morse Code - 12", "Complicated Wires - 13", "Wire Sequences - 14", "Mazes - 15", "Passwords - 16", "Venting Gas - 18", "Capacitor - 19", "Knobs - 20", "Appendix - 21"]
	$scope.text = "Index: \n" + pages.map(function(s){return "    "+s}).join("\n") + "\n\nred - \nblue - \nblack - \n";
}])
