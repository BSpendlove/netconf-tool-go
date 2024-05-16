/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/BSpendlove/netconf-tool-go/netconf-tool/cmd"
	_ "github.com/BSpendlove/netconf-tool-go/netconf-tool/cmd/operations"
	_ "github.com/BSpendlove/netconf-tool-go/netconf-tool/cmd/subscription"
)

func main() {
	cmd.Execute()
}
