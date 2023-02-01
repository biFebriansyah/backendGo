package command

import (
	"github.com/biFebriansyah/backintro/database/orm"
	"github.com/biFebriansyah/backintro/libs/server"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple backend golang",
	Long:  `golang backend with gorila/mux`,
}

func init() {
	initCommand.AddCommand(server.ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
	initCommand.AddCommand(orm.SeedCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}
