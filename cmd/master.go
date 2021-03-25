package main

import (
	"fmt"
	"os"

	"../env"
	"../node"
	"github.com/akamensky/argparse"
)

// Master Node
var masterNode node.Master

func main() {
	parser := argparse.NewParser("GPC Master", "Golang Process Commander")

	runCmd := parser.NewCommand("run", "Run Service")
	listCmd := parser.NewCommand("list", "[master] show slave resource list")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if runCmd.Happened() {
		node.MasterMode(masterNode, env.Ip, env.PortM)
	} else if listCmd.Happened() {
		fmt.Println("List Process")
	}

	pid := os.Getpid()
	parentpid := os.Getppid()
	fmt.Printf("The parent process id of %v is %v\n", pid, parentpid)
}
