var defaultPage = "home";

var myapp = new angular.module("myapp", ["ngTouch","ngResource", 'ngRoute', "ngSanitize"]);


// Sets up default page to be login and redirects every other one to 
myapp.config(function($routeProvider, $locationProvider) {
  $routeProvider

    .when('/404', {
      templateUrl : '/pages/404/404.html',
      controller : "404Ctl"
    })

    .when('/home', {
      templateUrl : '/pages/home/home.html',
      controller : "HomeCtl"
    })

    // .when('/about', {
    //   templateUrl : '/pages/about/about.html',
    //   controller : "HomeCtl"
    // })

    .when('/contact', {
      templateUrl : '/pages/contact/contact.html',
      controller : "HomeCtl"
    })

    // causes no path to go to default page
    .when('', {
      redirectTo : function () {
        console.log("empty")
        return "/" + defaultPage
      }
    })
    .when('/', {
      redirectTo : function () {
        console.log("empty")
        return "/" + defaultPage
      }
    })


    // causes unrecognized path to go to default page
    // redirects to login if not logged in
    .otherwise({
      redirectTo : function (a, b, c) {
        return "/404#"+b;
      }
    })

    // $locationProvider.html5Mode(true);

});


// enables attribute ng-enter which calls a function when enter is pressed
myapp.directive('ngEnter', function () {
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
myapp.controller("MainCtl", ["$scope", "$rootScope", "$resource", "$location",  function($scope, $rootScope, $resource, $location){

  $rootScope.info = "";
  $rootScope.error = "";
  $rootScope.infoDiv = "";
  $rootScope.firstHit = true;
  $rootScope.oldPage = "";

  
  $rootScope.LOREMIPSUM = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas hendrerit dignissim nisi vel gravida. Aenean at sem lobortis, efficitur leo nec, feugiat sem. Etiam a justo non nulla mattis tincidunt. Proin lobortis ipsum quis arcu suscipit, non efficitur leo cursus. Sed non sollicitudin arcu. Donec dapibus ultrices lacus, ullamcorper varius turpis vehicula quis. Mauris id felis ex. In vitae ante vulputate, viverra leo sed, luctus orci. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Curabitur a lacus ut orci semper dapibus id eget lorem. Fusce dignissim ipsum orci, malesuada imperdiet nisi maximus at. Fusce ac risus eget leo porta mattis. Praesent iaculis felis scelerisque, sollicitudin tellus placerat, ornare urna. Proin eu interdum neque. <br><br> Nam gravida elit a porta mattis. Phasellus placerat tincidunt est at facilisis. Phasellus a condimentum lectus. Aliquam interdum felis et libero gravida, luctus sagittis orci pellentesque. Vivamus suscipit vitae dolor sit amet dapibus. Quisque ullamcorper orci vel tellus pellentesque, in convallis arcu imperdiet. Nulla vulputate nisl eget interdum dignissim. Proin vitae dictum tellus. Morbi quis augue at ligula elementum congue quis ac risus. In sem odio, volutpat vitae arcu ornare, dapibus faucibus nunc. Mauris at commodo enim. Maecenas vel turpis felis. Ut sed arcu ac felis convallis elementum. Cras eu erat diam. Praesent sit amet felis nec ligula congue lacinia et ac justo. Nullam sit amet condimentum tortor. <br><br> Phasellus nulla turpis, porta vitae ligula a, convallis porttitor lorem. Donec ac luctus erat. Pellentesque vitae arcu vel ligula gravida aliquet quis vel turpis. Maecenas vestibulum eros vitae erat aliquam tempus. Donec consequat nulla varius dui semper tristique. Duis nec imperdiet risus. Nulla nisi arcu, suscipit maximus egestas a, rhoncus ut erat. Vestibulum venenatis posuere fringilla. Pellentesque dapibus eros eget neque feugiat viverra. Proin eu lectus tortor. Aenean bibendum tempor sapien nec hendrerit. Quisque magna sem, mollis vel lorem in, euismod fermentum tortor. <br><br> Ut ut sapien eget eros facilisis egestas vel sed magna. Sed nibh ex, commodo efficitur sapien a, scelerisque tempor massa. Mauris consequat suscipit ipsum, sit amet hendrerit tellus lobortis lobortis. Sed id dapibus massa. Nullam diam velit, tincidunt id mollis sed, ultricies in mauris. Mauris dignissim bibendum imperdiet. Mauris hendrerit vulputate volutpat. Curabitur in auctor augue, eu imperdiet nisl. <br><br> Pellentesque lacinia nulla tellus, et interdum odio sagittis nec. Praesent turpis metus, vestibulum et sagittis non, tincidunt at ipsum. Phasellus est purus, congue at mauris at, suscipit pulvinar ipsum. Phasellus vitae quam consequat, pellentesque neque quis, molestie felis. Morbi dictum eleifend purus. Pellentesque ultrices lobortis urna, at aliquam arcu efficitur nec. Phasellus vestibulum sem nunc, aliquam iaculis nibh venenatis ac. <br><br> Fusce non elit a orci semper gravida. Vivamus dapibus enim urna, eget congue libero scelerisque sit amet. Fusce ultrices erat eu urna luctus lacinia. Aliquam varius laoreet velit ac sollicitudin. Phasellus molestie eros augue, ac pellentesque nibh tincidunt at. Phasellus vitae dictum augue, a dictum diam. Phasellus egestas mattis elit ac egestas. Integer eu erat ac enim cursus eleifend nec quis elit. Fusce in sollicitudin nulla. Pellentesque ut feugiat mi, vel sodales est. Suspendisse accumsan purus ut lacus gravida egestas. Mauris pulvinar quam vitae magna sollicitudin mattis. <br><br> Nam risus lorem, pulvinar id finibus aliquet, condimentum a lorem. Ut id efficitur magna. Vestibulum tincidunt urna mauris, a convallis nulla dignissim eu. Mauris vel leo eros. Nullam in felis vestibulum, semper odio id, ornare est. Mauris quis quam nec sapien sollicitudin fermentum. Nam laoreet purus nec augue auctor condimentum. <br><br> Integer suscipit lobortis nisl id mattis. Cras mi nulla, vestibulum a erat blandit, facilisis aliquam odio. Vestibulum ornare lorem vitae lacinia viverra. Donec porta nisl vel justo bibendum pharetra. Pellentesque ut congue elit. Pellentesque posuere, magna vitae hendrerit hendrerit, tortor ipsum dapibus tellus, sit amet porta ante diam ut lorem. Donec vestibulum imperdiet velit. <br><br> Curabitur vitae orci vulputate, pellentesque felis ut, tincidunt mi. Quisque sagittis iaculis quam ac vulputate. Phasellus a tempus justo. Vestibulum lorem sem, feugiat quis purus et, efficitur ullamcorper elit. Pellentesque vestibulum sem augue, vel gravida risus sagittis a. Nunc ut purus a sapien dignissim vehicula. Nunc rhoncus turpis ac nunc maximus imperdiet. Praesent gravida velit nec nisl luctus, vitae bibendum metus commodo. <br><br> Nullam ac eleifend urna. Fusce eget gravida dolor. Quisque a metus faucibus, dignissim tortor sit amet, sodales leo. Pellentesque faucibus dui mi, vitae ornare metus pharetra non. Vestibulum semper nulla ut arcu finibus tristique. Curabitur malesuada orci sem, quis placerat mauris lacinia ac. Aliquam congue risus sed nulla ultrices egestas. Phasellus eleifend lorem eget massa cursus, pulvinar tincidunt ligula molestie. Donec et dolor turpis. Donec gravida metus quis ipsum volutpat, id auctor sem dignissim. Phasellus elementum ac orci tempus dignissim. Nullam malesuada, leo at mollis tristique, lacus mauris semper quam, quis porttitor ex nibh nec velit. Duis pretium mi sed dolor tristique pulvinar. Nam sit amet nisi ac nunc accumsan egestas."





  // does nothing if on a page besides login
  // if on login, sends to default page
  $rootScope.SendToPage = function() {
    page = $rootScope.oldPage;
    hash = $rootScope.oldHash;
    if ($rootScope.page == "login") {
      $rootScope.page = page;
      $location.path(page);
      $rootScope.info = "";
      // console.log("reset in sendtopage")
      $location.hash(hash);
    }

  }



  // returns html bold tezt
  $rootScope.boldText = function(text) {
    return "<strong>" + text + "</strong>";
  }
  

  // returns path==viewlocation
  $rootScope.isActive = function (viewLocation) { 
    return viewLocation === $location.path();
  }

  // returns true if there is a user in local storage
  // user is cleared whenever logged out
  $rootScope.loggedIn = function () {
    return localStorage.getItem('User') !== null;
  }

  // returns input if logged in
  // used to disable the links
  $rootScope.ifLoggedIn = function (str){
    if ($rootScope.loggedIn()) return str;
    return "";
  }

 

}]);
