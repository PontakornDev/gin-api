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
	route.GET("/getProduct", GetProduct)
	route.GET("/getAllProduct", GetAllProduct)
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

func GetProduct(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
  
	c.JSON(http.StatusOK, gin.H{"data": books})
  }