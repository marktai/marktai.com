let defaultPage = () => window.location.hostname === '52.11.108.65' ? 'index' : 'home';

var marktai = new angular.module("marktai", ["ngTouch","ngResource", 'ngRoute', "ngSanitize", 'ngFileUpload', "ngStorage"]);


// Sets up default page to be login and redirects every other one to 
marktai.config(function($routeProvider, $locationProvider) {
  $routeProvider .when('/404', {
      templateUrl : './pages/404/404.html',
      controller : "404Ctl"
    })

    .when('/home', {
      templateUrl : './pages/home/home.html',
      controller : "HomeCtl"
    })

    .when('/posts', {
      templateUrl : './pages/posts/posts.html',
      controller : "PostsCtl"
    })

    .when('/about', {
         templateUrl : './pages/about/about.html',
         controller : "AboutCtl"
    })

    .when('/contact', {
      templateUrl : './pages/contact/contact.html',
      controller : "ContactCtl"
    })

    .when('/corgis', {
      templateUrl : './pages/corgis/corgis.html',
      controller : "CorgisCtl"
    })

    .when('/bomb2/password', {
      templateUrl : './pages/password/password.html',
      controller : "PasswordCtl"
    })

    .when('/bom3/morse', {
      templateUrl : './pages/morse/morse.html',
      controller : "MorseCtl"
    })

    .when('/bomb', {
      templateUrl : './pages/bomb/bomb.html',
      controller : "BombCtl"
    })

    .when('/d', {
      templateUrl : './pages/decrypto/decrypto.html',
      controller : "DecryptoCtl"
    })

	.when('/shortlink', {
      templateUrl : './pages/shortlink/shortlink.html',
      controller : "ShortlinkCtl"
    })

	.when('/upload', {
      templateUrl : './pages/upload/upload.html',
      controller : "UploadCtl"
    })

    // .when('/where_is', {
    //   templateUrl : './pages/where_is/where_is.html',
    //   controller : "WhereIsCtl"
    // })

    .when('/fischer', {
      templateUrl : './pages/fischer/fischer.html',
      controller : "FischerCtl"
    })

    .when('/menu', {
      templateUrl : './pages/menu/menu.html',
      controller : "MenuCtl"
    })

    .when('/mootropolis', {
      templateUrl : './pages/mootropolis/mootropolis.html',
      controller : "MootropolisCtl"
    })

    .when('/interviewing', {
      templateUrl : './pages/interviewing/interviewing.html',
      controller : "InterviewingCtl"
    })

    .when('/diving', {
      templateUrl : './pages/diving/diving.html',
      controller : "DivingCtl"
    })

    .when('/index', {
      templateUrl : './pages/markieverse/markieverse.html',
      controller : "MarkieverseCtl"
    })

    .when('/blank_slate', {
      templateUrl : './pages/blank_slate/blank_slate.html',
      controller : "BlankSlateCtl"
    })

    .when('/blankSlate', {
      templateUrl : './pages/blank_slate/blank_slate.html',
      controller : "BlankSlateCtl"
    })

    // old url
	.when('/login', {
			
      controller : function() {
        window.location.replace('/meta-tic-tac-toe');
      },
      template : "<div></div>"
	})

    // causes no path to go to default page
    .when('', {
      redirectTo : function () {
        return "/" + defaultPage();
      }
    })
    .when('/', {
      redirectTo : function () {
        return "/" + defaultPage();
      }
    })


    // causes unrecognized path to go to default page
    // redirects to login if not logged in
    .otherwise({
      redirectTo : function (a, b, c) {
        return "/404#/#"+b;
      }
    })

    // $locationProvider.html5Mode(true);

});


// enables attribute ng-enter which calls a function when enter is pressed
marktai.directive('ngEnter', function () {
    return function (scope, element, attrs) {
        element.bind("keydown keypress", function (event) {
            if(event.which === 13) {
                scope.$apply(function (){
                    scope.$eval(attrs.ngEnter);
                });
 
                event.preventDefault();
            }
        });
    };
});

// inject the $resource dependency here
marktai.controller("MainCtl", ["$scope", "$rootScope", "$resource", "$location",  function($scope, $rootScope, $resource, $location){

  $rootScope.info = "";
  $rootScope.error = "";
  $rootScope.infoDiv = "";
  $rootScope.firstHit = true;
  $rootScope.oldPage = "";

  if ($location.absUrl().indexOf("/beta") != -1) {
     $rootScope.apiLocation = "/apiBeta"
  } else {

	$rootScope.apiLocation = "/api"
  }
}]);
