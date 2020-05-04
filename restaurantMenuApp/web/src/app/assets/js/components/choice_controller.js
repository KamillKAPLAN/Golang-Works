app.controller('choiceController', ['$scope', 'choiceService', '$routeParams', function($scope, choiceService, $routeParams) {
  $scope.choices = choiceService.choiceResource.query({
    categoryId: $routeParams.categoryId,
    productId: $scope.product.id,
    optionId: $scope.option.id
  });
}]);
