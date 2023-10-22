package commands

import (
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/migrate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrationCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate schema in database",
	Long:  "Need to write smth here",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := migrate.NewMigrateSession(&cfg.Database)
		if err != nil {
			log.Fatal(err)
		}
		err = s.Migrate()
		if err != nil {
			log.Fatal(err)
		}
	},
}
