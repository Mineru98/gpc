package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/jaypipes/ghw"
)

func cpuInfo() {
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	fmt.Printf("%v\n", cpu)

	for _, proc := range cpu.Processors {
		fmt.Printf(" %v\n", proc)
		for _, core := range proc.Cores {
			fmt.Printf("  %v\n", core)
		}
		if len(proc.Capabilities) > 0 {
			// pretty-print the (large) block of capability strings into rows
			// of 6 capability strings
			rows := int(math.Ceil(float64(len(proc.Capabilities)) / float64(6)))
			for row := 1; row < rows; row = row + 1 {
				rowStart := (row * 6) - 1
				rowEnd := int(math.Min(float64(rowStart+6), float64(len(proc.Capabilities))))
				rowElems := proc.Capabilities[rowStart:rowEnd]
				capStr := strings.Join(rowElems, " ")
				if row == 1 {
					fmt.Printf("  capabilities: [%s\n", capStr)
				} else if rowEnd < len(proc.Capabilities) {
					fmt.Printf("                 %s\n", capStr)
				} else {
					fmt.Printf("                 %s]\n", capStr)
				}
			}
		}
	}
}

func main() {
	parser := argparse.NewParser("GPC", "Golang Process Commander")
	modeArg := parser.String("m", "master", &argparse.Options{Required: true, Help: "You can choose between 'Master' and 'Slave'. ('Master', 'master', 'm' / 'Slave', 'slave', 's')"})
	pingArg := flag.String("p", "xxx.xxx.xxx.xxx", "Ping Test")
	numbPtr := flag.Int("numb", 42, "an int")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	fmt.Println(*parser)
}
