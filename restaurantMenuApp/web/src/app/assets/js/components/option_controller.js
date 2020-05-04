app.controller('optionController', ['$scope', '$localStorage', 'optionService', 'choiceService', '$routeParams',
function($scope, $localStorage, optionService, choiceService, $routeParams) {
  $scope.options = [];
  optionService.optionResource.query({
    categoryId: $routeParams.categoryId,
    productId: $scope.product.id
  }, function(options) {
    options.forEach(function(option, index) {
      choiceService.choiceResource.query({
        categoryId: $routeParams.categoryId,
        productId: $scope.product.id,
        optionId: option.id
      }, function(choices) {
        choices.forEach(function(choice) {
          choice.added = false;
        });
        option.choices = choices;
        $scope.options.push(option);
      });
    });
  });
  // Cart Start Code

  $scope.totalPrice = $scope.product.price;
  $scope.selected_choices = [];

  $scope.addChoiceToArray = function(choice) {
    $scope.selected_choices.push(choice);
    $scope.totalPrice += choice.price;
  }

  $scope.removeChoiceFromArray = function(choice) {
    for (var i = 0; i < $scope.selected_choices.length; i++) {
      if ($scope.selected_choices[i].id == choice.id) {
        $scope.selected_choices.splice(i, 1);
        $scope.totalPrice -= choice.price;
      }
    }
  }

  $scope.complete = function() {
    $scope.selected_options = [];
    $localStorage.cart.totalAmount += $scope.totalPrice;
    $scope.selected_choices.forEach(function(selected_choice) {
      $scope.options.forEach(function(option) {
        if (selected_choice.option_id == option.id) {
          let isIncluded = false
          $scope.selected_options.forEach(function(selected_option){
            if (selected_option.id == option.id) {
              isIncluded = true
            }
          })
          if (!isIncluded) {
            $scope.selected_options.push(option);
            // push yap
          }
        }
      });
    });

    $localStorage.cart.products.push($scope.product);
    $scope.close();
  }
}]);
