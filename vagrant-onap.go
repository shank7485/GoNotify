package main

import (
	//"flag"
	arg_funcs "github.com/shank7485/vagrant-onap-cli/arg_parser"
	"os"
	//"fmt"
	"flag"
)


func main()  {

	if len(os.Args) < 2 {
		arg_funcs.PrintInvalid()

	}

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)

	createTextPtr := createCommand.String("component", "", "Component name. (Required)")
	createBuildContainerPtr := createCommand.Bool("build", false, "Option to build container")
	createRunContainerPtr := createCommand.Bool("run", false, "Option to run container")


	switch os.Args[1] {
	case "help":
		arg_funcs.PrintBaseHelp()
	case "list":
		arg_funcs.PrintAvailableONAPComponents()
	case "create":
		createCommand.Parse(os.Args[2:])
	case "delete":
		//
	default:
		arg_funcs.PrintInvalid()
	}
}

func createCommand(commandType flag)  {
	if countCommand.Parsed() {

	}
	switch commandType {
	case createTextPtr:
		// do
	case createBuildContainerPtr:
		// do
	case createRunContainerPtr:
		// do
	}
}
