marktai.controller("FischerCtl", ["$scope", "$rootScope", "$interval", "$location", "$q", function($scope, $rootScope, $interval, $location, $q) {
	$rootScope.info = ""
	$rootScope.error = ""
	$rootScope.page = "fischer";


	var PERIOD = 100 // milliseconds
	$scope.fischer_time = 60 // seconds
	$scope.initial_time = 180 // seconds
	$scope.complement_time = 180 * 1000 // seconds

	$scope.timers = {};

	$scope.active_timer = null;

	$scope.timer_to_add = "";
	$scope.add_timer = function(){
		if ($scope.timer_to_add === '') return;
		$scope.timers[$scope.timer_to_add] = parseInt($scope.initial_time) * 1000;
		$scope.timer_to_add = '';
	};

	$scope.decrement_timer = function(n){
		if ($scope.active_timer != null) {
			$scope.timers[$scope.active_timer] -= n;
		}
	};

	$scope.complement_timer = function(timer){
		$scope.timers[timer] = $scope.complement_time - $scope.timers[timer];
	};

	$scope.switch_timer = function(new_timer, fischer_time){
		if (typeof fischer_time === "undefined") {
			fischer_time = parseInt($scope.fischer_time) * 1000;
		}
		if (new_timer == $scope.active_timer){
			return;
		};

		$scope.timers[new_timer] += fischer_time;
		$scope.active_timer = new_timer;
	};

	$interval(
		function(){
			$scope.decrement_timer(PERIOD);
		},
		PERIOD,
	);
}])
