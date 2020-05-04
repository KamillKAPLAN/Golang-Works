package ana

import (
	"github.com/GO/Hafta2/1OneDay/03/nine"
	"github.com/gin-gonic/gin"
)

func Ana(api *gin.RouterGroup) {
	api.DELETE("/ana", nine.Nine)
}
