myapp.controller("LoginCtl", ["$scope", "$rootScope", "$http", "$location", "$sce",  function($scope, $rootScope, $http, $location, $sce) {
	$rootScope.page = "login";
	
	$scope.usr = '';
	$scope.pwd = '';

	$scope.out = '';
	$scope.id = 0;
	$scope.secret = '';


	var login = function(user, pass) {
		var creds = {"User":user, "Password":pass};
		$http.post('/T9/login', creds).then(function(result){ 
			console.log(result);
			var data = result.data;
			$scope.id = data['UserID'];
			$scope.secret = data['Secret'];
			$scope.out = "ID: " + $scope.id + "| Secret: " + $scope.secret;
		}, function(error){
			$scope.out = "Login failed.";
		});
	}

	var verifySecret = function(user, secret) {
		var creds = {"User":user, "Secret":secret};
		$http.post('/T9/verifySecret', creds).then(function(result){ 
			console.log(result);
			var data = result.data;
			$scope.id = data['UserID'];
			$scope.secret = data['Secret'];
			$scope.out = "Verified| ID: " + $scope.id + "| Secret: " + $scope.secret;
		}, function(error){
			$scope.out = "Verification failed.";
		});
	}


	$scope.login = function() {
		login($scope.usr, $scope.pwd);
		$scope.pwd = '';
	}

	$scope.verifySecret = function() {
		verifySecret($scope.usr, $scope.secret);
	}
}])
