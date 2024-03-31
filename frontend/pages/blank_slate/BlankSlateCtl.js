marktai.controller("BlankSlateCtl", ["$scope", "$rootScope", "$interval", "$location", "$q", function($scope, $rootScope, $interval, $location, $q) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "blank_slate";

	$scope.prompts = ["AFTER____", "AIR____", "AMERICAN____", "ANY____", "APPLE____", "ARM____", "BABY____", "BACK____", "BAKED____", "BANANA____", "BATH____", "BEACH____", "BED____", "BEST____", "BIRTHDAY____", "BLACK____", "BLUE____", "BOOK____", "BOWLING____", "BREAK____", "BROKEN____", "BUS____", "CANDY____", "CAR____", "CASH____", "CHECK____", "CHOCOLATE____", "CHRISTMAS____", "COFFEE____", "DAY____", "DESK____", "DINNER____", "DOG____", "DOUBLE____", "DRIVE____", "DRUG____", "EVEN____", "EYE____", "FACE____", "FALSE____", "FINAL____", "FIRE____", "FIRST____", "FOR____", "GAS____", "GET____", "GIFT____", "GOLDEN____", "GOLF____", "GOOD____", "GRILLED____", "HAIR____", "HAND____", "HEAD____", "HEART____", "HEAVY____", "HIGH____", "HOME____", "HONEY____", "HOT____", "ICE____", "JET____", "JUNK____", "JUST____", "KICK____", "LAKE____", "LAST____", "LIFE____", "LOVE____", "LUNCH____", "MAGIC____", "MEAT____", "MICHAEL____", "MILK____", "MOTHER____", "MOTOR____", "NEVER____", "NOSE____", "ODD____", "OPEN____", "OUT____", "OVER____", "PARKING____", "PARTY____", "POST____", "RAIN____", "REST____", "RIGHT____", "RUBBER____", "SALT____", "SHOPPING____", "SHORT____", "SHOW____", "SHUT____", "SNOW____", "STUFFED____", "SUB____", "SUMMER____", "SUN____", "TENNIS____", "TRAFFIC____", "UNDER____", "WEDDING____", "WHAT'S____", "WHEEL____", "____ACHE", "____BAG", "____BALL", "____BALL", "____BASKET", "____BEAN", "____BEAR", "____BERRY", "____BIRD", "____BITE", "____BOAT", "____BONE", "____BOWL", "____BOY", "____BREAD", "____BREAD", "____BUG", "____BURN", "____CAGE", "____CAKE", "____CAN", "____CARDS", "____CASE", "____CHIP", "____CLASS", "____CLUB", "____COAT", "____CONTROL", "____COURT", "____CUP", "____DANCE", "____DATE", "____DAY", "____DECORATIONS", "____DOLL", "____DOOR", "____DOWN", "____DRIVER", "____FATHER", "____FIGHT", "____FINGER", "____FISH", "____FLAKES", "____FOOD", "____FRAME", "____FREE", "____FRIEND", "____GLASS", "____GLOVE", "____GUN", "____HOLE", "____HOUR", "____HOUSE", "____HOUSE", "____JAM", "____JUICE", "____KEEPER", "____LAND", "____LANGUAGE", "____LIGHT", "____LUCK", "____MARKET", "____MATE", "____MOUTH", "____NAME", "____PAD", "____PAD", "____PAPER", "____PIE", "____POLE", "____POTATO", "____RING", "____ROOM", "____ROOM", "____SALAD", "____SAUCE", "____SCOPE", "____SEAT", "____SEAT", "____SHARK", "____SHIP", "____SHOP", "____SHOT", "____SHOWER", "____SITE", "____STAR", "____STATION", "____STICK", "____STORM", "____TABLE", "____TICKET", "____TIME", "____TOAST", "____TOWN", "____TREE", "____WALK", "____WASH", "____WATER", "____WAY", "____WEIGHT", "____WOMAN", "____WORD", "____WORK", "____YARD"];

	$scope.prompts = $scope.prompts.map(value => ({ value, sort: Math.random() }))
    .sort((a, b) => a.sort - b.sort)
    .map(({ value }) => value);

	$scope.prompt_index = 0;

	$scope.players = {};

	$scope.player_to_add = "";
	$scope.add_player = function(){
		if ($scope.player_to_add === '') return;
		$scope.players[$scope.player_to_add] = 0;
		$scope.player_to_add = '';
	};

	$scope.update_player = function(player, delta){
		$scope.players[player] = Number($scope.players[player]) + delta;
	};

	$scope.update_prompt_index = function(delta){
	    $scope.prompt_index += delta;
	}
}])
