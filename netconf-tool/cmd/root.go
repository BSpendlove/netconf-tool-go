/*
Copyright © 2023 Brandon Spendlove <brandonspendlove@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "netconf-tool",
	Short: "CLI Tool to interact with NETCONF Servers",
	Long: `netconf-tool is a Go based port of my original Python CLI tool netconf-tool
which uses the click module.

The idea of this tool is to provide some basic NETCONF functionality on demand so you can
interact with a NETCONF server on the fly with basic operations instead of having to write
temporary code for example to gather a NETCONF subscription and event data, or grab the running
configuration to then store and parse locally for offline development.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
