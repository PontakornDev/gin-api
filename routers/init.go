package router

import (
	"net/http"
	"pontakorn322/demo_gin_api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "demo api")
	})

	return router
}

func SetUpRouteGroup(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	controllers.ProductEndPoint(v1)
}
