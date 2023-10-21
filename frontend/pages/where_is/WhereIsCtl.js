marktai.controller("WhereIsCtl", ["$scope", "$rootScope", "$route", "$window", function($scope, $rootScope, $route, $window) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "where_is";
	$scope.show_circle = false;

	$scope.no_circle_url = 'https://www.marktai.com/download/54689/IMG-8565.JPG';
  $scope.circle_url = 'https://www.marktai.com/download/54689/IMG-8565%20copy.JPG';

	$scope.image_url = function(){
		if (!$scope.show_circle){
			return $scope.no_circle_url;
		}
		else {
			return $scope.circle_url;
		}
	}

	$scope.toggle_circle = function(){
		$scope.show_circle = !$scope.show_circle;
	}
}])
