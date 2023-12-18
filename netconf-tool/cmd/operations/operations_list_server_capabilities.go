/*
Copyright © 2023 Brandon Spendlove <brandonspendlove@gmail.com>
*/
package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BSpendlove/netconf-tool-go/netconf-tool/utils"
	"github.com/spf13/cobra"
)

var (
	printCli   bool
	exportCli  bool
	exportJson string
)

// listServerCapabilitiesCmd represents the listServerCapabilities command
var listServerCapabilitiesCmd = &cobra.Command{
	Use:   "list-server-capabilities",
	Short: "Print and/or Export NETCONF Server Capabilities",
	Long: `Gather NETCONF Server Capabilities and either print them to stdout using
a string format or JSON format.

You can also export the full URL which includes the scheme, path, queries, fragments etc..
to a JSON file`,
	Run: func(cmd *cobra.Command, args []string) {
		// Don't allow CLI export and print, no point in this
		if printCli && exportCli {
			log.Fatal("printcli and exportcli arguments shouldn't be set together, pick one or the other")
		}

		sessionArgs := utils.BuildNetconfArgs(cmd)
		sessionConfig := utils.BuildSSHConfig(&sessionArgs)

		// Setup NETCONF Session and Context
		session := utils.SetupNetconfSession(sessionArgs, &sessionConfig)
		ctx := context.Background()

		// Get NETCONF Server capabilities
		serverCapabilities, err := utils.BuildCapabilitiesMap(session.ServerCapabilities())
		if err != nil {
			log.Fatal(err)
		}

		if err := session.Close(ctx); err != nil {
			log.Print(err)
		}

		// Print NETCONF Server Capabilities to stdout
		if printCli {
			for i, s := range *serverCapabilities {
				fmt.Printf("(%d) %s\n", i, s.String())
			}
		}

		d, err := json.MarshalIndent(serverCapabilities, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		// Export NETCONF Server Capabilities to stdout
		if exportCli {
			fmt.Println(string(d))
		}

		// Export NETCONF Server Capabilities to JSON file
		if exportJson != "" {
			if !strings.HasSuffix(exportJson, ".json") {
				exportJson = exportJson + ".json"
			}
			err := os.WriteFile(exportJson, d, 0644)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("exported %d NETCONF server capabilities to %s", len(*serverCapabilities), exportJson)
		}
	},
}

func init() {
	listServerCapabilitiesCmd.Flags().BoolVar(&printCli, "printcli", false, "Print the NETCONF Server Capabilities as a string to stdout")
	listServerCapabilitiesCmd.Flags().BoolVar(&exportCli, "exportcli", false, "Export each Capability as a string to stdout")
	listServerCapabilitiesCmd.Flags().StringVar(&exportJson, "exportjson", "", "Filename to export NETCONF Server Capabilities to")

	operationsCmd.AddCommand(listServerCapabilitiesCmd)
}
