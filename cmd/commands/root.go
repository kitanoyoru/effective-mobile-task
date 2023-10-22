package commands

import (
	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfg config.Config

var devFlag bool

var rootCmd = &cobra.Command{
	Use:   "effective-mobile-task",
	Short: "bla bla bla",
	Long:  "more bla bla bla",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&devFlag, "dev", false, "use development version")
}

func initConfig() {
	if err := config.UnmarshalFromEnv(&cfg, &config.UnmarshalOptions{
		Dev: devFlag,
	}); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(serverCommand)
	rootCmd.AddCommand(migrationCommand)
}
