/*
Copyright © 2023 Brandon Spendlove <brandonspendlove@gmail.com>
*/
package operations

import (
	"context"
	"fmt"
	"log"

	"github.com/BSpendlove/netconf-tool-go/netconf-tool/utils"
	"github.com/spf13/cobra"
)

// getConfigCmd will attempt to use the NETCONF <get-config/> operation and return in either XML or JSON
var getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Run a <get-config/> NETCONF Operation against a NETCONF Server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Setup Args and SSH ClientConfig
		sessionArgs := utils.BuildNetconfArgs(cmd)
		sessionConfig := utils.BuildSSHConfig(&sessionArgs)

		// yang.

		// Setup NETCONF Session and Context
		session := utils.SetupNetconfSession(sessionArgs, &sessionConfig)
		ctx := context.Background()

		// Send NETCONF Command
		getConfig, err := session.GetConfig(ctx, "running")
		if err != nil {
			log.Fatal(err)
		}

		getConfigStr := string(getConfig)

		fmt.Println(getConfigStr)

		if err := session.Close(ctx); err != nil {
			log.Print(err)
		}
	},
}

func init() {
	operationsCmd.AddCommand(getConfigCmd)
}
