myapp.controller("LoginCtl", ["$scope", "$rootScope", "$http", "$location", "$sce", "$q",  function($scope, $rootScope, $http, $location, $sce, $q) {
	$rootScope.page = "login";
	
	$scope.usr = '';
	$scope.pwd = '';

	$scope.out = '';
	$scope.id = 0;
	$scope.secret = '';


	$scope.game = $location.hash()

	$scope.gameData = {}
    $scope.board = ""
    $scope.boardArray = []

    $scope.players = []
    $scope.boxes = []
    $scope.squares = []



    for (var i = 0; i < 9; i++) {
		var box = {
			'Owned' : 0,
			'Squares': []
		};
    	for (var j = 0; j < 9; j++) {
    		box['Squares'].push(0);
    	}
    	$scope.boardArray.push(box)
    	$scope.boxes.push('' + i)
    	$scope.squares.push('' + i)
    }

    $scope.player = ''
    $scope.box = $scope.boxes[0]
    $scope.square = $scope.squares[0]

        var xImg;
        var oImg; 


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
		$http.get('/T9/games/' + gameID + '/board').then(function(result){ 
			$scope.boardArray = result.data["Board"]
			for (var i = 0; i < 9; i++){
				for (var j = 0; j < 9; j++) {
					$scope.loadIcon(i, j);
				}
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


	$scope.getGame = function() {
		getGame($scope.game);
	}

    $scope.makeMove = function(game, player, box, square) {
    	makeMove($scope.game, $scope.player, $scope.box, $scope.square);
    }

    $scope.canvasClicked = function(a, b) {
    	console.log(a + ", " + b);
		$scope.boardArray[a].Squares[b] += 1
		if ($scope.boardArray[a].Squares[b] > 2) {
			$scope.boardArray[a].Squares[b] = 0
		}
		$scope.loadIcon(a, b);
    }

    $scope.loadIcon = function (a, b) {
		var canvas  = document.getElementById("box"+a+"-"+b);
		var context = canvas.getContext("2d");
		console.log(a + ", " + b + "=" + $scope.boardArray[a].Squares[b]);
    	if ($scope.boardArray[a].Squares[b] == 0) {
			context.clearRect(0, 0, canvas.width, canvas.height);
		} else if ($scope.boardArray[a].Squares[b] == 1) {
			xImg.then(function(image){
				context.drawImage(image, 0, 0);
			}, function(error){
				console.log("x failed to load")
			})
		} else if ($scope.boardArray[a].Squares[b] == 2) {
			oImg.then(function(image){
				context.drawImage(image, 0, 0);
			}, function(error){
				console.log("o failed to load")
			})
		}
    }

    function loadImage(src) {
        return $q(function(resolve,reject) {
            var image = new Image();
            image.src = src;
            image.onload = function() {
              console.log("loaded image: "+src);
              resolve(image);
            };
            image.onerror = function(e) {
              reject(e);
            };
        })
    }

	$scope.getGame();

    xImg = loadImage("https://www.marktai.com/img/x.png");
    
    oImg = loadImage("https://www.marktai.com/img/o.jpg");

}])
