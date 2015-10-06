myapp.controller("PostsCtl", ["$scope", "$rootScope", "$http", "$location", function($scope, $rootScope, $http, $location) {
	$rootScope.page = "posts";

	$scope.post = $location.hash()

	$scope.postData = {}

	var getPost = function(post) {
		$http.get('/api/posts/' + post).then(function(result){ 
			$scope.postData = result.data["Post"]
		})
	}

	$scope.generateTag = function() {
		var retStr = "Posted"
		if ($scope.postData["Author"] !== "") {
			retStr += " by " + $scope.postData["Author"]
		}
		if ($scope.postData["Created"] !== "") {
			retStr += " on " + new Date($scope.postData["Created"]).toDateString()
		}
		if (retStr == "Posted") {
			return ""
		}
		return retStr
	}

	getPost($scope.post)


	console.log($scope.postData)
}])
