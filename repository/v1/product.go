package repository

import (
	"pontakorn322/demo_gin_api/models"
	"pontakorn322/demo_gin_api/repository"
)

type ReqProduct struct {
	Code  string  `json:"code"`
	Price float32 `json:"price"`
}

func AddProduct(r ReqProduct) error {
	data := models.Product{
		Code:  r.Code,
		Price: r.Price,
	}
	if err := repository.DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProduct()([]models.Product , error) {
	var products []models.Product
	if err := repository.DB.Find(&products).Error; err != nil {
		return nil,err
	}
	return products,nil
}

func GetProduct(id int)(models.Product , error) {
	var products models.Product
	if err := repository.DB.First(&products,"id = ?",id).Error; err != nil {
		return products,err
	}
	return products,nil
}

func DeleteProduct(id int) (models.Product , error) {
	var products models.Product
	if err := repository.DB.Delete(&products,"id = ?",id).Error; err != nil {
		return products,err
	}
	return products,nil
}

func UpdateProduct(r models.Product) (models.Product , error) {
	var products models.Product
	if err := repository.DB.Updates(&r).Error; err != nil {
		return products,err
	}
	return products,nil
}
