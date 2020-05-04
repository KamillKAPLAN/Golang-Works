app.controller('adminController', ['$scope', 'categoryService', '$modal', function($scope, categoryService, $modal) {

  // $scope.category = categoryService.category.get({
  //   categoryId: '1'
  // });
  // console.log($scope.category);

  $scope.categories = categoryService.categoryResource.query();

  $scope.new_category = {
      name : '',
      list_order: 5
  }
  $scope.addItem = function () {
    $scope.errortext = "";
    if ($scope.new_category.name != '') {
      new categoryService.categoryResource($scope.new_category).$save(function(createdCategory){
        $scope.categories.push(createdCategory);
      });
      $scope.new_category.name = '';
    } else {
      $scope.errortext = "Boş bırakmayın a";
    }
  }


  $scope.removeItem = function(i, c) {
    categoryService.categoryResource.delete({
      categoryId: c.id
    });
    $scope.categories.splice(i, 1);
  }

  $scope.update_category = {};
  $scope.showCategoryModal = function(category) {
    $scope.update_category = category;
    $scope.showModal();
  }

  $scope.updateCategory = function(c) {
    new categoryService.categoryResource($scope.update_category).$update({
      categoryId: c.id
    },function(updatedCategory){
      c = updatedCategory;
      myOtherModal.hide();
        // window.location.reload();
    });
  }
  var myOtherModal = $modal({
    scope: $scope,
    placement: 'center',
    templateUrl: 'assets/tpl/modals/category-modal.html',
    show: false
  });
  $scope.showModal = function() {
    myOtherModal.$promise.then(myOtherModal.show);
  };

  // $scope.categories = [];
  // $http({
  //   method : 'GET',
  //   url    : 'http://localhost:1323/categories'
  // }).then(function mySuccess(response) {
  //   $scope.categories = response.data;
  //   console.log($scope.categories);
  // }, function myError(err) {
  //     $scope.error = err.statusText;
  //   });
}])
