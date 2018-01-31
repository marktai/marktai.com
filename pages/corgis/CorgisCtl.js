marktai.controller("CorgisCtl", ["$scope", "$rootScope", "$route", "$window", function($scope, $rootScope, $route, $window) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "corgis";

	var nonce = (new Date()).getTime();

	var generate_nonce = function(){
		return nonce += 1;
	};

	$scope.generate_corgi_image = function(){
		return "http://corgis.marktai.com/image?nonce=" + generate_nonce().toString();
	};

	$scope.generate_corgi_gif = function(){
		return "http://corgis.marktai.com/gif?nonce=" + generate_nonce().toString();
	};

	$scope.refresh_page = function(){
		$window.scrollTo(0, 0);
		$route.reload();
	};
	
}])
