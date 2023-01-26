package models

import "time"

type User struct {
	ID        uint    `gorm:"primary_key; auto_increment; index;" json:"productID"`
	FirstName string  `grom:"not null" json:"firstname"`
	LastName  float32 `grom:"not null" json:"lastname"`
	Age       int     `grom:"not null" json:"age"`

	CreatedAt *time.Time `gorm:"DEFAULT:now();" json:"-"`
	UpdatedAt *time.Time `gorm:"DEFAULT:now()" json:"-"`
}
