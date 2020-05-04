package nine

import (
	"net/http"

	"github.com/GO/Hafta2/1OneDay/03/model"
	"github.com/gin-gonic/gin"
)

var nine [5]model.Nine

func Nine(c *gin.Context) {
	for i := range nine {
		nine[i].Gorev = "silindi"
	}
	c.JSON(http.StatusOK, nine)
}
