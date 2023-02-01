package orm

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "for running database migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", true, "for running auto migration")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "for running auto reset migration")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := NewDb()
	if err != nil {
		return err
	}

	if migDown {
		return db.Migrator().DropTable(&models.Product{}, &models.Cart{}, &models.CartItem{}, &models.User{})
	}

	if migUp {
		return db.AutoMigrate(&models.Product{}, &models.Cart{}, &models.CartItem{}, &models.User{})
	}

	return nil

}
