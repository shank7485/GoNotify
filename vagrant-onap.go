package main

import (
	"flag"
	"fmt"
	funcs "github.com/shank7485/vagrant-onap-cli/arg_parser"
	"os"
)

func main() {

	funcs.Init()

	if len(os.Args) < 2 {
		funcs.PrintInvalid()

	}

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	createDircPtr := createCommand.String("d", "", "Directory where the Vagrant File is locatied. (Required)")
	createTextPtr := createCommand.String("component", "", "Component name. (Required)")
	createBuildContainerPtr := createCommand.Bool("build", false, "Option to build container")
	createRunContainerPtr := createCommand.Bool("run", false, "Option to run container")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteDircPtr := deleteCommand.String("d", "", "Directory where the Vagrant File is locatied. (Required)")
	deleteTextPtr := deleteCommand.String("component", "", "Component name. (Required)")

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	listSupported := listCommand.Bool("supported", false, "List of supported operations")

	switch os.Args[1] {
	case "help":
		funcs.PrintBaseHelp()
	case "list":
		listCommand.Parse(os.Args[2:])
		//funcs.PrintAvailableONAPComponents()
	case "create":
		createCommand.Parse(os.Args[2:])
	case "delete":
		deleteCommand.Parse(os.Args[2:])
	default:
		funcs.PrintInvalid()
	}

	if createCommand.Parsed() {
		if *createDircPtr == "" || *createTextPtr == "" || funcs.Services_map[*createTextPtr] == false {
			fmt.Printf(
				"Vagrant Directory %s or Service %s not found\n", *createDircPtr, *createTextPtr)
			funcs.PrintInvalid()
		}

		if *createBuildContainerPtr == false &&
			*createRunContainerPtr == false {
			funcs.CreateONAPComponent(*createDircPtr, *createTextPtr)
		}

		if *createBuildContainerPtr == true {
			funcs.CreateONAPcomponent_DockerBuild(*createDircPtr, *createTextPtr)
		}

		if *createRunContainerPtr == true {
			funcs.CreateONAPComponent_DockerRun(*createDircPtr, *createTextPtr)
		}

	}

	if deleteCommand.Parsed() {
		if *deleteDircPtr == "" || *deleteTextPtr == "" || funcs.Services_map[*deleteTextPtr] == false {
			fmt.Printf(
				"Vagrant Directory %s or Service %s not found\n", *deleteDircPtr, *deleteTextPtr)
			funcs.PrintInvalid()
		}

		funcs.DeleteONAPComponent(*deleteDircPtr, *deleteTextPtr)
	}

	if listCommand.Parsed() {
		if *listSupported == true {
			funcs.ListONAPComponents()
		} else {
			funcs.CheckRunningONAPComponents()
		}

	}
}
