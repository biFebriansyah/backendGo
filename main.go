package main

import (
	"log"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/biFebriansyah/backintro/command"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
