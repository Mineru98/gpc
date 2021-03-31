package main

import (
	"fmt"
	"os"

	"../env"
	"../node"
	"github.com/akamensky/argparse"
)

// Slave
var slaveNode node.Slave

func main() {
	parser := argparse.NewParser("GPC Slave", "Golang Process Commander")

	runCmd := parser.NewCommand("run", "Run Service")
	addCmd := parser.NewCommand("add", "[slave] add CLI process")
	listCmd := parser.NewCommand("list", "[slave] show process list")

	modeIpAddressArg := runCmd.String("a", "address",
		&argparse.Options{
			Required: true,
			Help:     "Setting Master Ip Address",
		})
	modePortArg := runCmd.String("p", "port",
		&argparse.Options{
			Required: true,
			Help:     "Setting Master Port",
		})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if runCmd.Happened() {
		node.SlaveMode(slaveNode, env.Ip, env.Port, *modeIpAddressArg, *modePortArg)
	} else if addCmd.Happened() {
		fmt.Println("Add Process")
	} else if listCmd.Happened() {
		fmt.Println("List Process")
	}

	pid := os.Getpid()
	parentpid := os.Getppid()
	fmt.Printf("The parent process id of %v is %v\n", pid, parentpid)

}
