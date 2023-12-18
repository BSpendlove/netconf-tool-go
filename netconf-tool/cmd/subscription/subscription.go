/*
Copyright © 2023 Brandon Spendlove <brandonspendlove@gmail.com>
*/
package subscription

import (
	"fmt"

	"github.com/BSpendlove/netconf-tool-go/netconf-tool/cmd"
	"github.com/spf13/cobra"
)

// subscriptionCmd represents the subscription command
var subscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Short: "Perform various NETCONF Subscription commands",
	Long: `Perform various NETCONF Subscribe Operations against a NETCONF Server to listen
to events and create subscription filters to gather NETCONF alarms/alerts/events`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Peform a NETCONF Subscription Operations by using one of the various sub-commands.")
	},
}

func init() {
	cmd.RootCmd.AddCommand(subscriptionCmd)
}
