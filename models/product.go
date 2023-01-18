package models

import "time"

type Product struct {
	ID    uint    `gorm:"primary_key; auto_increment; index;" json:"productID"`
	Code  string  `grom:"not null" json:"code"`
	Price float32 `grom:"not null" json:"price"`

	CreatedAt *time.Time `gorm:"DEFAULT:now();" json:"-"`
	UpdatedAt *time.Time `gorm:"DEFAULT:now()" json:"-"`
}
