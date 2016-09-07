marktai.controller("UploadCtl", ["$scope", "$rootScope", "$http", "$location", "$q", "Upload", "LoginService", function($scope, $rootScope, $http, $location, $q, Upload, LoginService) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "home";

	$scope.files = [];
	$scope.doneFiles = [];
	$scope.totalFiles = [];
	$scope.fileProgress = {};
	$scope.fileLinks = {};

	$scope.username = "";
	$scope.password = "";
	$scope.stayLoggedIn = false;

	$scope.userID = -1;
	$scope.secret = "";

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
		if ($scope.userID > 0) {
			return "./download/" + $scope.userID + "/" + file.name;
		}
		return "./download/" + file.name;
	}

	$scope.setCreds = function(creds) {
		$scope.userID = creds['UserID'];
		$scope.secret = creds['Secret'];
		$scope.username = creds['Username'];
	};

	$scope.login = function() {
		LoginService.login($scope.username, $scope.password, $scope.stayLoggedIn).then(function(creds) {
			$scope.setCreds(creds);
		}, function(error) {
			console.log(error);
		});
		$scope.password = "";
	}

    $scope.upload = function () {
        if ($scope.totalFiles && $scope.totalFiles.length) {
            for (var i = 0; i < $scope.totalFiles.length; i++) {
			  var file = $scope.totalFiles[i];
              if (!file.$error) {
				  var urlWithoutRoot = "/upload";
				  var authHeaders = LoginService.genAuthHeaders(urlWithoutRoot, $scope.secret);
				  authHeaders['UserID'] = $scope.userID;
				  var _ = function(file) {
					$scope.fileProgress[file.name] = "0%";
					Upload.upload({
						url: $rootScope.apiLocation + urlWithoutRoot,
						data: {
						  file: file, 
						},
						headers: authHeaders,
					}).then(function (resp) {
						var index = $scope.totalFiles.indexOf(file);
						if (index >= 0) {
							$scope.totalFiles.splice(index, 1);
						}
						delete $scope.fileProgress[file];
						$scope.fileLinks[file.name] = generateFileLink(file);
						$scope.doneFiles.push(file);

					}, function(error){
						$scope.fileProgress[file.name] = "Error: " + error.data.Error;
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
	};
	
	$scope.init = function() {
		LoginService.checkLocalStorageLogin().then(function(creds) {
			$scope.setCreds(creds);
		}, function(error) {
			console.log(error);
		});
	};

	$scope.init();

}])
