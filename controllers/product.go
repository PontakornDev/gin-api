package controllers

import (
	"net/http"
	repo "pontakorn322/demo_gin_api/repository/v1"
	"pontakorn322/demo_gin_api/utils"
	"strconv"
	"github.com/gin-gonic/gin"
)

func ProductEndPoint(router *gin.RouterGroup) {
	route := router.Group("/product")
	route.POST("/addProduct", InsertProduct)
	route.GET("/getProduct/:id", GetProduct)
	route.GET("/getAllProduct", GetAllProduct)
	route.DELETE("/deleteProduct/:id", DeleteProduct)
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

func GetAllProduct(ctx *gin.Context) {
	res, err := repo.GetAllProduct(); 
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       "Getall Product",
		Description: "Getall Success",
		Item: res,
	}))
}

func GetProduct(ctx *gin.Context) { 
	paramID := ctx.Param("id") 
	productID, err:=strconv.Atoi(paramID)
	res, err := repo.GetProduct(productID); 
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       "Get Product",
		Description: "Get Product Success",
		Item: res,
	}))
}


func DeleteProduct(ctx *gin.Context) { 
	paramID := ctx.Param("id") 
	productID, err:=strconv.Atoi(paramID)
	res, err := repo.DeleteProduct(productID); 
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       "Delete Product",
		Description: "Delete Product Success",
		Item: res,
	}))
}

