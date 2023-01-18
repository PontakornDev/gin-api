package controllers

import (
	"net/http"
	repo "pontakorn322/demo_gin_api/repository/v1"
	"pontakorn322/demo_gin_api/utils"

	"github.com/gin-gonic/gin"
)

func ProductEndPoint(router *gin.RouterGroup) {
	route := router.Group("/product")
	route.POST("/addProduct", InsertProduct)
}

func InsertProduct(ctx *gin.Context) {
	var req repo.ReqProduct
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if err := repo.AddProduct(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       "Add Product",
		Description: "Add Success",
	}))
}
