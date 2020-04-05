package internal

import (
	"github.com/gin-gonic/gin"
)

func StartHTTPServer() {
	r := gin.Default()

	r.GET("/date", GetDate)

	r.POST("/date", SetDate)

	_ = r.Run()
}
