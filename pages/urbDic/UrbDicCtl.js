marktai.controller("UrbDicCtl", ["$scope", "$rootScope", "$http", "$location", "$q", function($scope, $rootScope, $http, $location, $q) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "home";


	$scope.term = $location.hash()
	$scope.results = [];

	var lookupTerm = function(term) {
		return $http.get('/urban_dictionary/v0/define?term=' + term).then(function(result){ 
			var data = result.data;
			if (data["result_type"] == "no_results") {
				return $q.reject("No results found");
			}
			return $q.resolve(data);
		}, function(error){
			return $q.reject(error);
		});			
	}

	$scope.lookupTerm = function() {
		$scope.results = [];
		$rootScope.error = "";
		lookupTerm($scope.term).then(function(data){
			console.log(data);
			$scope.results = data.list;
		}, function(error){
			$rootScope.error = error;
		})	

		$location.hash($scope.term);
	}

	if ($scope.term != "") {
		$scope.lookupTerm();
	}
}])
