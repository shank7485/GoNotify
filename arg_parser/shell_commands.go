package arg_parser

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

type fn func(*bufio.Scanner)

var BaseHelp struct {
	help        string
	list        string
	createClone string
	createBuild string
	createRun   string
	delete      string
}

var Services_map = make(map[string]bool)

func Init() {
	InitServiceMap()
	InitPrintBaseHelp()
}

func InitPrintBaseHelp() {
	BaseHelp.help = "vagrant-onap help\tHelp\n"
	BaseHelp.list = "vagrant-onap list\tList of available Vagrant ONAP components.\n"
	BaseHelp.createClone = "vagrant-onap create -d <VagrantFile path> -component=<name>\tCreate component to only clone repos.\n"
	BaseHelp.createBuild = "vagrant-onap create -d <VagrantFile path> -component=<name> --build\tCreate component by cloning and building containers.\n"
	BaseHelp.createRun = "vagrant-onap create -d <VagrantFile path> -component=<name> --run\tCreate component by cloning, building and running containers.\n"
	BaseHelp.delete = "vagrant-onap delete -d <VagrantFile path> -component=<name>\tDelete component\n"
}

func InitServiceMap() {
	Services_map["sdnc"] = true
	Services_map["appc"] = true
	Services_map["ccsdk"] = true
	Services_map["policy"] = true
	Services_map["portal"] = true
	Services_map["mr"] = true
	Services_map["vfc"] = true
	Services_map["multicloud"] = true
}

func PrintInvalid() {
	fmt.Println("Invalid input\n")
	fmt.Println("Try 'vagrant-onap help' for help.")
	os.Exit(1)
}

func PrintBaseHelp() {
	fmt.Println(BaseHelp.help)
	fmt.Println(BaseHelp.list)
	fmt.Println(BaseHelp.createClone)
	fmt.Println(BaseHelp.createBuild)
	fmt.Println(BaseHelp.createRun)
	fmt.Println(BaseHelp.delete)
}

func PrintAvailableONAPComponents() {
	for k := range Services_map {
		fmt.Printf("%s\n", k)
	}
}

func PrintGeneric(scanner *bufio.Scanner) {
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), ",")
		fmt.Printf("%s\n", result[len(result)-1])
	}
}

func PrintRunningList(scanner *bufio.Scanner) {
	output := []string{}
	for scanner.Scan() {

		result := strings.Split(scanner.Text(), ",")
		value := result[len(result)-1]
		ok := Services_map[strings.TrimSpace(value)]

		if ok {
			output = append(output, strings.TrimSpace(value))
		}
	}

	if len(output) != 0 {
		for val := range output {
			fmt.Println(val)
		}
	} else {
		fmt.Println("No VMs are running.")
	}
}

func RunShell(cmdName string, cmdArgs []string, f fn) error {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()

	scanner := bufio.NewScanner(cmdReader)
	var stderr bytes.Buffer

	if err != nil {
		fmt.Printf("Error creating StdoutPipe for Cmd. Error: %s", err)
		os.Exit(1)
	}

	go f(scanner)

	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting Cmd. Error: %s", err)
		os.Exit(1)
	}

	// To handle Control C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\nStopping.\n")
		cmd.Process.Kill()
		os.Exit(1)
	}()

	cmd.Stderr = &stderr

	err = cmd.Wait()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + stderr.String())
	}

	return err
}
