/*
Copyright © 2023 Brandon Spendlove <brandonspendlove@gmail.com>
*/
package operations

import (
	"fmt"

	"github.com/BSpendlove/netconf-tool-go/netconf-tool/cmd"
	"github.com/spf13/cobra"
)

var (
	host          string // NETCONF Server IP
	port          int    // NETCONF Server Port
	timeout       int    // SSH Timeout
	username      string // Username used in SSH Authentication
	password      string // Password used in SSH Authentication
	pubkey        string // Path of public key if key auth is preferred over password auth
	pubkeypass    string // Private key password if key auth is preferred over password auth
	hostkeyignore bool   // Ignore HostKeys
)

// operationsCmd represents the operations command
var operationsCmd = &cobra.Command{
	Use:   "operations",
	Short: "Perform various NETCONF Operation commands",
	Long: `Perform various NETCONF Operations against a NETCONF Server such as
<get> and <get-config>.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Peform a NETCONF Operation by using one of the various sub-commands.")
	},
}

func init() {
	cmd.RootCmd.AddCommand(operationsCmd)

	// Persistent Flags used in all sub `operations` commands.
	operationsCmd.PersistentFlags().StringVar(&host, "host", "", "NETCONF Server Host")
	operationsCmd.PersistentFlags().IntVar(&port, "port", 830, "NETCONF Server Port")
	operationsCmd.PersistentFlags().IntVar(&timeout, "timeout", 5, "SSH Timeout")
	operationsCmd.PersistentFlags().StringVar(&username, "username", "", "NETCONF Server Username for Authentication")
	operationsCmd.PersistentFlags().StringVar(&password, "password", "", "NETCONF Server Password for Authentication")
	operationsCmd.PersistentFlags().StringVar(&pubkey, "pubkey", "", "SSH Key if public key authentication is preferred")
	operationsCmd.PersistentFlags().StringVar(&pubkeypass, "pubkeypass", "", "SSH Key Password if public key authentication is preferred")
	operationsCmd.PersistentFlags().BoolVar(&hostkeyignore, "hostkeyignore", true, "Ignore Hostkeys")

	// Required Arguments
	operationsCmd.MarkPersistentFlagRequired("host")
	operationsCmd.MarkPersistentFlagRequired("username")
}
