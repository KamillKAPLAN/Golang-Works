app.controller('categoryController', ['$scope', '$localStorage', 'categoryService', function($scope, $localStorage, categoryService) {
  $scope.categories = categoryService.categoryResource.query();

  var categoiess = $scope.categories;
  $scope.searchCategory = {}
  $scope.convert = "";
  $scope.$watch("searchCategory.name", function(newValue, oldValue) {
    if ( newValue !== oldValue) {
      $scope.convert = "";
      var degis = $scope.searchCategory.name;
      for (var i = 0; i < categoiess.length; i++) {
        if (categoiess[i].name.toLowerCase().includes(degis.toLowerCase())) {
          $scope.convert = "var";
          break;
        } else {
          $scope.convert = "Yok";
        }
      }
      console.log(oldValue + ":" + newValue);
    }
  }, true);

}]);
