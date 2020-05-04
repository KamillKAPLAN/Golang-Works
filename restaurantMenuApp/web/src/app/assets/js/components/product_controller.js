app.controller('productController', ['$scope', '$localStorage', 'productService', 'categoryService', '$routeParams', '$modal',
function($scope, $localStorage, productService, categoryService, $routeParams,$modal) {
  $scope.products = productService.productResource.query({
    categoryId: $routeParams.categoryId
  });
  $scope.category = categoryService.categoryResource.get({
    categoryId: $routeParams.categoryId
  });

  $scope.product = {};
  $scope.showProductModal = function(product) {
    $scope.product = product;
    $scope.showModal();
  }
  $scope.optionModal = $modal({
    scope: $scope,
    placement: 'center',
    templateUrl: 'assets/tpl/modals/option-modal.html',
    show: false
  });
  $scope.showModal = function() {
    $scope.optionModal.$promise.then($scope.optionModal.show);
  };

  $scope.close = function() {
    $scope.optionModal.hide();
  }

  //  Cart Start Code
  // prod = {
  //   name:'asd',
  //   price: 5
  //   options: [{}, {
  //     name: 'aşsdl',
  //     choices: [{}, {}]
  //   }]
  // }

}]);
