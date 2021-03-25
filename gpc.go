package main

import (
	"fmt"
	"os"

	"./env"
	"./node"
	"./utils/network"
	"github.com/akamensky/argparse"
)

var masterNode node.Master

var slaveNode node.Slave

func main() {
	parser := argparse.NewParser("gpc", "Golang Process Commander")

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
	modeIpAddressArg := runCmd.String("a", "address",
		&argparse.Options{
			Required: false,
			Help:     "Setting Master Ip Address",
			Default:  env.Ip,
		})
	modePortArg := runCmd.String("p", "port",
		&argparse.Options{
			Required: false,
			Help:     "Setting Master Port",
			Default:  env.PortM,
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
	pingPeriodArg := pingCmd.Int("P", "period",
		&argparse.Options{
			Required: false,
			Help:     "Period of check in ms",
			Default:  1000,
		})
	pingLimitCountArg := pingCmd.Int("n", "count",
		&argparse.Options{
			Required: false,
			Help:     "Ping Count",
			Default:  4,
		})

	slaveListArg := listCmd.Flag("s", "slave",
		&argparse.Options{
			Required: false,
			Help:     "Print Apply Process List from Slave",
		})
	masterListArg := listCmd.Flag("m", "master",
		&argparse.Options{
			Required: false,
			Help:     "Print Apply Slave List from Master",
		})

	versionArg := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:     "Check Version",
		})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if runCmd.Happened() {
		pid := os.Getpid()
		parentpid := os.Getppid()
		fmt.Printf("The parent process id of %v is %v\n", pid, parentpid)
		switch *modeArg {
		case "Master", "master", "M", "m":
			node.MasterMode(masterNode, *modeIpAddressArg, *modePortArg)
		case "Slave", "slave", "S", "s":
			node.SlaveMode(slaveNode, *modeIpAddressArg, *modePortArg)
		}
	} else if stopCmd.Happened() {
		fmt.Println("Stop Service")
	} else if pingCmd.Happened() {
		network.Ping(*pingAddressArg, *pingPortArg, *pingTimeoutArg, *pingPeriodArg, *pingLimitCountArg)
	} else if addCmd.Happened() {
		fmt.Println("Add Process")
	} else if listCmd.Happened() {
		fmt.Println("List Process")
		if *masterListArg == true && *slaveListArg == false {
			fmt.Println("m")
		} else if *masterListArg == false && *slaveListArg == true {
			fmt.Println("s")
		} else {
			fmt.Println("no")
		}
	} else if *versionArg {
		fmt.Println(env.BuildVersion + "." + env.BuildCode)
	}
}
