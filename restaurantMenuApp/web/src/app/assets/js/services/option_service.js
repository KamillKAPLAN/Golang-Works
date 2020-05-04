app.factory('optionService', ['$resource', function($resource) {
  var optionResource = $resource(
    'http://localhost:1323/categories/:categoryId/products/:productId/options/:optionId',
    {
      categoryId: '@id',
      productId: '@id',
      optionId: '@id'      
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
    optionResource: optionResource
  }

}]);
