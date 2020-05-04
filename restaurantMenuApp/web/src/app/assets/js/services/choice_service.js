app.factory('choiceService', ['$resource', function($resource) {
  var choiceResource = $resource(
    'http://localhost:1323/categories/:categoryId/products/:productId/options/:optionId/choices/:choiceId',
    {
      categoryId: '@id',
      productId: '@id',
      optionId: '@id',
      choiceId: '@id'
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
    choiceResource: choiceResource
  }

}]);
