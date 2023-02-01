package seeder

import "github.com/biFebriansyah/backintro/database/orm/models"

var ProductSeed = models.Products{
	models.Product{
		Name:        "indomie",
		Price:       "2000",
		Image:       "http://image.com",
		Description: "indomie original",
	},
	models.Product{
		Name:        "indomie ayam bawang",
		Price:       "2000",
		Image:       "http://image.com",
		Description: "indomie rasa ayam bawang",
	},
	models.Product{
		Name:        "indomie aceh",
		Price:       "2000",
		Image:       "http://image.com",
		Description: "indomie rasa mie khas aceh",
	},
	models.Product{
		Name:        "indomie nasi padang",
		Price:       "2000",
		Image:       "http://image.com",
		Description: "indomie rasa nasi padang",
	},
}
