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

	var getGame = function(gameID) {
		$http.get('/T9/games/' + gameID + '/string').then(function(result){ 
			$scope.board = ""
			var boardArray = result.data["Board"]
            for (line of boardArray) {
                $scope.board += line + "\n"
            }
		})
		$http.get('/T9/games/' + gameID).then(function(result){ 
			$scope.gameData = result.data["Game"]
    		$scope.players = $scope.gameData["Players"]
    		$scope.player = $scope.players[Math.floor($scope.gameData["Turn"]/10)] + ""
    		if ($scope.gameData["Turn"]%10 < 9) {
	    		$scope.box = $scope.gameData["Turn"]%10 + ""
    		}
		})
	}

    var makeMove = function(game, player, box, square) {

    	var time = Math.floor(Date.now() / 1000);
		var urlWithoutT9 = '/games/' + game + '?Player=' + player + '&Box=' + box + '&Square=' + square;
    	var message = time + ":" + urlWithoutT9;
		var hash = CryptoJS.HmacSHA256(message, $scope.secret);
		// var hashInBase64 = CryptoJS.enc.Base64.stringify(hash);


		var url = '/T9' + urlWithoutT9;
		var data = '';
		var config = {
			'headers': {
				'HMAC' : hash,
				'Encoding' : 'hex',
				'Time-Sent' : time,
			}
		}
        $http.post(url, data, config).then(function(result){
                $scope.getGame($scope.game)
        }, function(error) {
                $scope.error = error.data["Error"]
        })
    }

	$scope.login = function() {
		login($scope.usr, $scope.pwd);
		$scope.pwd = '';
	}

	$scope.verifySecret = function() {
		verifySecret($scope.usr, $scope.secret);
	}


	$scope.game = $location.hash()

	$scope.gameData = {}
    $scope.board = ""

    $scope.players = []
    $scope.boxes = []
    $scope.squares = []

    for (var i = 0; i < 9; i++) {
    	$scope.boxes.push('' + i)
    	$scope.squares.push('' + i)
    }

    $scope.player = ''
    $scope.box = $scope.boxes[0]
    $scope.square = $scope.squares[0]

	$scope.getGame = function() {
		getGame($scope.game);
	}

    $scope.makeMove = function(game, player, box, square) {
    	makeMove($scope.game, $scope.player, $scope.box, $scope.square);
    }

	$scope.getGame();
}])
