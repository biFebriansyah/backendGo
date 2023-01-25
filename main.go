package main

import (
	"log"
	"net/http"

	"github.com/biFebriansyah/backintro/router"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r, err := router.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("app runn on port 3001")
	err = http.ListenAndServe(":3001", r)
	if err != nil {
		log.Fatal(err)
	}
}
