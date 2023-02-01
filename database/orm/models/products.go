package models

import (
	"time"
)

type Product struct {
	ProductId   string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Name        string    `json:"name" valid:"type(string)"`
	Price       string    `json:"price" valid:"type(string)"`
	Image       string    `json:"image" valid:"-"`
	Description string    `json:"description" valid:"type(string)"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdateAt    time.Time `json:"update_at" valid:"-"`
}

type Products []Product

func (Product) TableName() string {
	return "products"
}
