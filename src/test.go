package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func test() {
	parser := argparse.NewParser("commands", "Test")

	startCmd := parser.NewCommand("start", "Start Process")

	stopCmd := parser.NewCommand("stop", "Stop Prcoess")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if startCmd.Happened() {
		fmt.Println("Start Process")
	} else if stopCmd.Happened() {
		fmt.Println("Stop Process")
	} else {
		err := fmt.Errorf("bag arguments")
		fmt.Print(parser.Usage(err))
	}
}
