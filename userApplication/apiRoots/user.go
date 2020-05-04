package apiRoots

import (
	"github.com/gin-gonic/gin"
	"github.com/PROJECTS/user_application_go/apiControllers"
)

func UserApi(api *gin.RouterGroup) {

	api.GET("/users",apiControllers.GetUserList)
	api.GET("/user/:id", apiControllers.GetUserListById)

	// api.POST("/user", apiControllers.PostUser)

	api.PUT("/user", apiControllers.PutUser)

	api.DELETE("/user/:id", apiControllers.DeleteUser)
}