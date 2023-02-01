package seeder

import "github.com/biFebriansyah/backintro/database/orm/models"

var UserSeed = models.Users{
	{
		Username: "users",
		Email:    "users@email.com",
		Role:     "user",
		Password: "$2a$10$n3ljave8ZvvSwOGZ40VK3OunbAvpIuZHDPj88317MbW6D70GYz4xy",
	},
	{
		Username: "admin",
		Email:    "admin@email.com",
		Role:     "admin",
		Password: "$2a$10$r6nunPtknjyDs4IkCo9Jpec.o5TP0lIRhY1PL6UcjXNjQfrS6H3H2",
	},
	{
		Username: "users2",
		Email:    "users2@email.com",
		Role:     "user",
		Password: "$2a$10$72Y4sMeISN5qaEKvDsSwqedZfCHHOv69BzDAZhdhwowPQ.KMEvSgy",
	},
}
