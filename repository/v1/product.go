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

func GetProduct()([]models.Product , error) {
	var products []models.Product
	if err := repository.DB.Find(&products).Where("is_active = true").Error; err != nil {
		return nil,err
	}
	return products,nil
}

