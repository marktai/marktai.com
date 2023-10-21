marktai.controller("DecryptoCtl", ["$scope", "$rootScope", "$route", "$window", "$location", function($scope, $rootScope, $route, $window, $location) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "decrypto";
	
	$scope.error = ""
	$scope.alphabet = 'abcdefghijklmnopqrstuvwxyz'.split('');

	$scope.word = "decrypto";
	$scope.index = 1;

	$scope.press_letter = function(l){
		if (l == $scope.word[$scope.index]){
			if (l == 'o'){
				window.location.replace('http://decrypto.marktai.com')
			}
			$scope.index += 1
			$location.hash($scope.word.substring(1, $scope.index))
		} else {
			$scope.error = "Not " + l
		}
	}

	$scope.populate_scope = function(){
		var cur = $location.hash()
		var i;
		for (i = 0; i < cur.length; i++){
			if($scope.word[i+1] != cur[i]){
				break;
			}
		}
		var word = cur.substring(0, i)
		$location.hash(word)
		$scope.index = word.length + 1
	}

	$scope.populate_scope()
	
}])
