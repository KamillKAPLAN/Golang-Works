package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Headline struct {
	EffectiveDate      string `json:"effectiveDate"`
	EffectiveEpochDate int64  `json:"EffectiveEpochDate"`
	Severity           int32  `json:"severity"`
}

func getHeadLine(c *gin.Context) {
	mode := Headline{
		EffectiveDate:      "asd",
		EffectiveEpochDate: 25,
		Severity:           23,
	}
	c.JSON(http.StatusOK, mode)
}

func main() {
	app := gin.Default()
	
	app.GET("/apiTest", getHeadLine)

	app.Run(":1478")
}
