package arg_parser

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"os/signal"
	"syscall"
	"bytes"
	"strings"
)

type fn func(*bufio.Scanner)

func PrintInvalid()  {
	fmt.Println("Invalid input\n")
	fmt.Println("Try 'vagrant-onap help' for help.")
	os.Exit(1)
}

func PrintBaseHelp(){
	fmt.Printf("Example commands:\n")
	fmt.Printf("vagrant-onap help\tHelp\n")
	fmt.Printf("vagrant-onap list\tList of available Vagrant ONAP components.\n")
	fmt.Printf(
		"vagrant-onap create -d <VagrantFile path> -component=<name>\tCreate component to only clone repos.\n")
	fmt.Printf(
		"vagrant-onap create -d <VagrantFile path> -component=<name> --build\tCreate component by cloning and building containers.\n")
	fmt.Printf(
		"vagrant-onap create -d <VagrantFile path> -component=<name> --run\tCreate component by cloning, building and running containers.\n")
	fmt.Printf(
		"vagrant-onap delete -d <VagrantFile path> -component=<name>\tDelete component\n")
}

func PrintAvailableONAPComponents()  {
	fmt.Printf("sdnc\n")
	fmt.Printf("appc\n")
	fmt.Printf("policy\n")
	fmt.Printf("portal\n")
	fmt.Printf("vfc\n")
}

func CheckRunningONAPComponents() {
	VagrantGlobalStatus()
}

func PrintGeneric(scanner *bufio.Scanner){
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), ",")
		fmt.Printf("%s\n", result[len(result)-1])
	}
}

func PrintRunningList(scanner *bufio.Scanner){
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), ",")
		value := result[len(result)-1]
		fmt.Printf("%s\n", value)
	}
}

func RunShell(cmdName string, cmdArgs []string, f fn) error {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()

	scanner := bufio.NewScanner(cmdReader)
	var stderr bytes.Buffer

	if err!=nil {
		fmt.Printf("Error creating StdoutPipe for Cmd. Error: %s", err)
		os.Exit(1)
	}

	go f(scanner)

	err = cmd.Start()
	if err!=nil {
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

	if err!=nil {
		fmt.Println(fmt.Sprint(err) + stderr.String())
	}

	return err
}
