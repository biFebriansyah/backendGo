package server

import (
	"log"
	"net/http"
	"os"

	"github.com/biFebriansyah/backintro/router"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "for running api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	mainRoute, err := router.NewApp()
	if err != nil {
		return err
	}

	var addrs string = "0.0.0.0:8080"
	if pr := os.Getenv("PORT"); pr != "" {
		addrs = ":" + pr
	}

	log.Println("apprun on port ", addrs)
	return http.ListenAndServe(addrs, mainRoute)
}
