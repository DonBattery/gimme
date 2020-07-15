package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (app *App) buildInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the Gimme folder at the GIMME_ROOT_DIR",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.initApp()
		},
	}
	return initCmd
}

func (app *App) initApp() error {
	conf := app.getConf()
	fmt.Printf("Viola:\n%+v\n", conf)
	return nil
}
