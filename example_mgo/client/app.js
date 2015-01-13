var myapp = angular.module("myapp", []);
  
myapp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/', {
        templateUrl: 'templates/catalogue.html',
        controller: 'ProductsController'
      }).
      when('/detail/:productId', {
        templateUrl: 'templates/detail.html',
        controller: 'DetailController'
      }).
      when('/cart', {
        templateUrl: 'templates/cart.html',
        controller: 'MyCartController'
      }).
      when('/checkout', {
        templateUrl: 'templates/checkout.html',
        controller: 'CheckoutController'
      }).
      otherwise({
        redirectTo: '/'
      });
  }]);

// inject the $resource dependency here
myapp.controller("ProductsController", ["$scope", function($scope){
	
}]);
myapp.controller("DetailController", ["$scope", function($scope){
	
}]);
myapp.controller("MyCartController", ["$scope", function($scope){
	
}]);
myapp.controller("CheckoutController", ["$scope", function($scope){
	
}]);