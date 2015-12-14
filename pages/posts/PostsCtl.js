myapp.controller("PostsCtl", ["$scope", "$rootScope", "$http", "$location", "$sce",  function($scope, $rootScope, $http, $location, $sce) {
	$rootScope.page = "posts";

	$scope.post = $location.hash()

	$scope.postData = {}

	var getPost = function(post) {
		$http.get('/api/posts/' + post).then(function(result){ 
			$scope.postData = result.data["Post"]
            for (var par in $scope.postData.Content) {
                $scope.pars.append($sce.trustAsHtml(par))
            }

		})
	}

    $scope.pars = []

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

}])
