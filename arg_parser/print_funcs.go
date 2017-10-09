package arg_parser

import (
	"fmt"
	"os"
)

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
		"vagrant-onap create -component=<name>\tCreate component to only clone repos.\n")
	fmt.Printf(
		"vagrant-onap create -component=<name> --build\tCreate component by cloning and building containers.\n")
	fmt.Printf(
		"vagrant-onap create -component=<name> --run\tCreate component by cloning, building and running containers.\n")
	fmt.Printf(
		"vagrant-onap delete -component=<name>\tDelete component\n")
}

func PrintAvailableONAPComponents()  {
	fmt.Printf("sdnc\n")
	fmt.Printf("appc\n")
	fmt.Printf("policy\n")
	fmt.Printf("portal\n")
	fmt.Printf("vfc\n")
	fmt.Printf("Example usage: vagrant-onap create -component=sdnc\n")
}



