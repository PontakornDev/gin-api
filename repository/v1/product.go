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
	if err := repository.DB.Create(&models.Product{
		Code:  r.Code,
		Price: r.Price,
	}).Error; err != nil {
		return err
	}
	return nil
}
