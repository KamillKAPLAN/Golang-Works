app.controller('mainController', ['$scope', '$localStorage', function($scope, $localStorage) {
  $scope.localStorage = $localStorage;
  $localStorage.loggedIn = false;

  $scope.pretendLogin = function() {
    $localStorage.cart = {
        products: [],
        totalAmount: 0
    }

    $localStorage.loggedIn = true;
  }

  $scope.pretendLogout = function() {
    $localStorage.$reset({
      loggedIn: false
    });
  }
}])
