myapp.controller("PostsCtl", ["$scope", "$rootScope", "$http", "$location", function($scope, $rootScope, $http, $location) {
	$rootScope.page = "posts";

	$scope.post = $location.hash()

	$scope.postData = {}

	var getPost = function(post) {
		$http.get('/api/posts/' + post).then(function(result){ 
			$scope.postData = result.data["Post"]
		})
	}

	getPost($scope.post)
}])
