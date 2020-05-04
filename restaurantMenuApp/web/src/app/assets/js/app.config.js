app
.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
    $routeProvider
    .when("/", {
        templateUrl : "assets/tpl/components/categories.html",
        // controller: 'categoryController'
    })
    .when("/categories/:categoryId", {
        templateUrl : "assets/tpl/components/products.html",
        // controller: 'productController'
    })
    .when("/categories/:categoryId/products", {
        templateUrl : "assets/tpl/components/options.html",
        // controller: 'optionController'
    })
    .when("/admin", {
        templateUrl : "assets/tpl/components/admin.html",
        // controller: 'adminController'
    })
    .otherwise({
			redirectTo: '/'
		});
    $locationProvider.hashPrefix('');
}])

.config(function($modalProvider) {
  angular.extend($modalProvider.defaults, {
    animation: 'am-flip-x',
    placement: 'center'
  });
})
