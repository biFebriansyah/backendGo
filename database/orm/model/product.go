package model

import "time"

type Product struct {
	ProductId uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "product"
}

type Products []Product
