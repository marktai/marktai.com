myapp.controller("HomeCtl", ["$scope", "$rootScope", "$resource", "$location", function($scope, $rootScope, $resource, $location) {
  $rootScope.info = ""
  $rootScope.error = ""
  $rootScope.page = "home";
}])