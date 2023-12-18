/*
Copyright © 2023 Brandon Spendlove <brandonspendlove@gmail.com>
*/
package subscription

import (
	"log"

	"github.com/spf13/cobra"
)

// localCmd represents the local command
var subLocalCmd = &cobra.Command{
	Use:   "local",
	Short: "Attempts to create a <create-subscription> event and interact locally",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("command currently not implemented")
	},
}

func init() {
	subscriptionCmd.AddCommand(subLocalCmd)
}
