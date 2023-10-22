package commands

import (
	"github.com/kitanoyoru/effective-mobile-task/internal/models"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/store"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  "Need to write smth here",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
