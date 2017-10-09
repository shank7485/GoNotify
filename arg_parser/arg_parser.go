package arg_parser

import (
	"fmt"
	"os"
)

func PrintInvalid()  {
	fmt.Println("Invalid input")
	PrintBaseHelp()
	os.Exit(1)
}

func PrintBaseHelp(){
	fmt.Printf("\n")
	fmt.Printf("Example commands:\n")
	fmt.Printf("vagrant-onap help\n")
	fmt.Printf("vagrant-onap list\n")
	fmt.Printf("vagrant-onap create -component=<name>\n")
	fmt.Printf("vagrant-onap create -component=<name> --build\n")
	fmt.Printf("vagrant-onap create -component=<name> --run\n")
	fmt.Printf("vagrant-onap delete -component=<name>\n")
}

func PrintAvailableONAPComponents()  {
	fmt.Printf("\n")
	fmt.Printf("sdnc\n")
	fmt.Printf("appc\n")
	fmt.Printf("policy\n")
	fmt.Printf("portal\n")
	fmt.Printf("vfc\n")
	fmt.Printf("\nExample usage: vagrant-onap create -component=sdnc\n")
}



