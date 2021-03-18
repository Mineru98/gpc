package main

import (
	"fmt"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("GPC", "Golang Process Commander")

	runCmd := parser.NewCommand("run", "Run Service")
	stopCmd := parser.NewCommand("stop", "Stop Service")
	pingCmd := parser.NewCommand("ping", "Ping Test. Be sure to check before connecting with the 'slaves'.")
	addCmd := parser.NewCommand("add", "[slave] add CLI process")
	listCmd := parser.NewCommand("list", "[master] show slave resource list / [slave] show process list")

	modeArg := runCmd.Selector("m", "mode",
		[]string{"Master", "master", "M", "m", "Slave", "slave", "S", "s"},
		&argparse.Options{
			Required: true,
			Help:     "select GCP operation mode",
		})

	pingAddressArg := pingCmd.String("a", "address",
		&argparse.Options{
			Required: true,
			Help:     "Test target ip address to check",
		})
	pingPortArg := pingCmd.Int("p", "port",
		&argparse.Options{
			Required: false,
			Help:     "Test target port to check",
			Default:  80,
		})
	pingTimeoutArg := pingCmd.Int("t", "timeout",
		&argparse.Options{
			Required: false,
			Help:     "Connection timeout in ms. Should be less than check 'period'",
			Default:  500,
		})
	periodArg := pingCmd.Int("P", "period",
		&argparse.Options{
			Required: false,
			Help:     "Period of check in ms",
			Default:  1000,
		})

	versionArg := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:     "Check Version",
		})

	fmt.Println(*versionArg)
	if runCmd.Happened() {
		fmt.Println(*modeArg)
		fmt.Println("Start Service")
	} else if stopCmd.Happened() {
		fmt.Println("Stop Service")
	} else if pingCmd.Happened() {
		fmt.Println("Add Process")
		fmt.Println(*pingAddressArg)
		fmt.Println(*pingPortArg)
		fmt.Println(*pingTimeoutArg)
		fmt.Println(*periodArg)
		// ping(pingAddressArg, pingPortArg, pingTimeoutArg, periodArg)
	} else if addCmd.Happened() {
		fmt.Println("Add Process")
	} else if listCmd.Happened() {
		fmt.Println("List Process")
	} else if *versionArg {
		fmt.Println("0.0.1")
	} else {
		err := fmt.Errorf("bag arguments")
		fmt.Print(parser.Usage(err))
	}
}
