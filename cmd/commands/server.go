package commands

import (
	"github.com/kitanoyoru/effective-mobile-task/internal/app"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  "Need to write smth here",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.NewApp(&cfg)
		if err != nil {
			log.Fatal(err)
		}

		if err := app.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
