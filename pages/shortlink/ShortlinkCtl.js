marktai.controller("Base3Ctl", ["$scope", "$rootScope", "$http", "$location", function($scope, $rootScope, $http, $location) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "home";

	$scope.decimal = "10";
	$scope.base3 = "101";


	var convertToDecimal = function(base3) {
		$http.get('/apiBeta/base3/decode?encoded=' + base3).then(function(result){
			$scope.decimal = result.data["Decoded"]
		}, function(error) {
			$scope.error = error;
		})
	}


	var convertToBase3 = function(decimal) {
		$http.get('/apiBeta/base3/encode?value=' + decimal).then(function(result){
			$scope.base3 = result.data["Encoded"]
		}, function(error) {
			$scope.error = error;
		})
	}

	$scope.convertToDecimal = function () {
		convertToDecimal($scope.base3)
	}

	$scope.convertToBase3 = function () {
		convertToBase3($scope.decimal)
	}
}])
