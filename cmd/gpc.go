package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"../env"
	"../utils/network"
	"github.com/akamensky/argparse"
)

// 동일한 프로세스 실행시 실행 중지
func checkDupleProcess(mode string) bool {
	switch mode {
	case "M", "m", "Master", "master":
		return false
	case "S", "s", "Slave", "slave":
		return false
	}
	return true
}

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

	debugArg := parser.Flag("d", "debug",
		&argparse.Options{
			Required: false,
			Help:     "Debugging mode",
			Default:  false,
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

	// 디버그 모드 활성화 / 비활성화
	env.Debug = *debugArg

	if runCmd.Happened() {
		if !checkDupleProcess(*modeArg) {
			switch *modeArg {
			case "Master", "master", "M", "m":
				if env.Debug {
					if runtime.GOOS == "Windows" {
						runCmd := exec.Command("go", "run", "cmd/master.go", "run")
						stdout, err := runCmd.StdoutPipe()

						err = runCmd.Start()
						if err != nil {
							log.Printf("Command finished with error: %v", err)
						}

						bResult, _ := ioutil.ReadAll(stdout)
						fmt.Println(string(bResult))
					} else {
						exec.Command("echo test")
						runCmd := exec.Command("nohup", "./master-cmd.sh", "run", ">", "/dev/null", "2>&1", "&")
						stdout, err := runCmd.StdoutPipe()

						err = runCmd.Start()
						if err != nil {
							log.Printf("Command finished with error: %v", err)
						}

						bResult, _ := ioutil.ReadAll(stdout)
						fmt.Println(string(bResult))
					}
				} else {
					if runtime.GOOS == "Windows" {
						runCmd := exec.Command("cmd.exe", "/c", "master.vbs", "run")
						stdout, err := runCmd.StdoutPipe()

						err = runCmd.Start()
						if err != nil {
							log.Printf("Command finished with error: %v", err)
						}

						bResult, _ := ioutil.ReadAll(stdout)
						fmt.Println(string(bResult))
					} else {
						runCmd := exec.Command("nohup", "./master", ">", "/dev/null", "2>&1", "&")
						stdout, err := runCmd.StdoutPipe()

						err = runCmd.Start()
						if err != nil {
							log.Printf("Command finished with error: %v", err)
						}

						bResult, _ := ioutil.ReadAll(stdout)
						fmt.Println(string(bResult))
					}
				}
			case "Slave", "slave", "S", "s":
				if env.Debug {
					if runtime.GOOS == "Windows" {
						runCmd := exec.Command("go", "run", "cmd/slave.go", "run", "-a", *modeIpAddressArg, "-p", *modePortArg)
						stdout, err := runCmd.StdoutPipe()

						err = runCmd.Start()
						if err != nil {
							log.Printf("Command finished with error: %v", err)
						}

						bResult, _ := ioutil.ReadAll(stdout)
						fmt.Println(string(bResult))
					} else {
						runCmd := exec.Command("nohup", "./slave-cmd.sh", "run", "-a", *modeIpAddressArg, "-p", *modePortArg, ">", "/dev/null", "2>&1", "&")
						stdout, err := runCmd.StdoutPipe()

						err = runCmd.Start()
						if err != nil {
							log.Printf("Command finished with error: %v", err)
						}

						bResult, _ := ioutil.ReadAll(stdout)
						fmt.Println(string(bResult))
					}
				} else {
					if runtime.GOOS == "Windows" {
						runCmd := exec.Command("cmd.exe", "/c", "slave.vbs", "run", "-a", *modeIpAddressArg, "-p", *modePortArg)
						stdout, err := runCmd.StdoutPipe()

						err = runCmd.Start()
						if err != nil {
							log.Printf("Command finished with error: %v", err)
						}

						bResult, _ := ioutil.ReadAll(stdout)
						fmt.Println(string(bResult))
					} else {
						runCmd := exec.Command("nohup", "./slave", "run", "-a", *modeIpAddressArg, "-p", *modePortArg, ">", "/dev/null", "2>&1", "&")
						stdout, err := runCmd.StdoutPipe()

						err = runCmd.Start()
						if err != nil {
							log.Printf("Command finished with error: %v", err)
						}

						bResult, _ := ioutil.ReadAll(stdout)
						fmt.Println(string(bResult))
					}
				}
			}
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
