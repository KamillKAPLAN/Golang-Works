app.factory('categoryService', ['$resource', function($resource) {
  var categoryResource = $resource(
    'http://localhost:1323/categories/:categoryId',
    {
      categoryId: '@id'
    },
    {
      'query': {
        method: 'GET',
        isArray: true
      },
      'update': {
        method: 'PUT'
      }
    }
  )

  return {
    categoryResource: categoryResource
  }

}]);
