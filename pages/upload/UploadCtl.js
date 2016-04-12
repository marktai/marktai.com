marktai.controller("UploadCtl", ["$scope", "$rootScope", "$http", "$location", "$q", "Upload", function($scope, $rootScope, $http, $location, $q, Upload) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "home";

	$scope.files = [];
	$scope.doneFiles = [];
	$scope.totalFiles = [];
	$scope.fileProgress = {};
	$scope.fileLinks = {};

	$scope.$watch('files', function() {
		if ($scope.files && $scope.files.length) {
			$scope.totalFiles = $scope.totalFiles.concat($scope.files);
			for (var i = 0; i < $scope.files.length; i++) {
				$scope.fileProgress[$scope.files[i].name] = "";
			}
			$scope.files = []
		}
	});

	var generateFileLink = function(file) {
		return "./upload/" + file.name;
	}	


    $scope.upload = function () {
        if ($scope.totalFiles && $scope.totalFiles.length) {
            for (var i = 0; i < $scope.totalFiles.length; i++) {
			  var file = $scope.totalFiles[i];
              if (!file.$error) {
				  var _ = function(file) {
					$scope.fileProgress[file.name] = "0%";
					Upload.upload({
						url: $rootScope.apiLocation + "/upload",
						data: {
						  file: file, 
						}
					}).then(function (resp) {
						var index = $scope.totalFiles.indexOf(file);
						if (index >= 0) {
							$scope.totalFiles.splice(index, 1);
						}
						delete $scope.fileProgress[file];
						$scope.fileLinks[file.name] = generateFileLink(file);
						$scope.doneFiles.push(file);

					}, function(error){
						console.log(error);
					}, function (evt) {
						var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
						$scope.fileProgress[file.name] = progressPercentage + "%";
					});
				  }

				  _(file);
              }
            }
        }
    };

	$scope.clear = function() {
		$scope.totalFiles = [];
		$scope.files = [];
	}
	

}])
