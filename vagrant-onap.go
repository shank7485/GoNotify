package main

import (
	funcs "github.com/shank7485/vagrant-onap-cli/arg_parser"
	"os"
	"flag"
	"fmt"
)

func main()  {

	var services_map = make(map[string]bool)

	services_map["sdnc"] = true
	services_map["appc"] = true
	services_map["policy"] = true
	services_map["portal"] = true
	services_map["vfc"] = true

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


	switch os.Args[1] {
	case "help":
		funcs.PrintBaseHelp()
	case "list":
		funcs.PrintAvailableONAPComponents()
	case "create":
		createCommand.Parse(os.Args[2:])
	case "delete":
		deleteCommand.Parse(os.Args[2:])
	default:
		funcs.PrintInvalid()
	}

	if createCommand.Parsed() {
		if *createDircPtr == "" || *createTextPtr == "" || services_map[*createTextPtr] == false {
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
		if *deleteDircPtr == "" || *deleteTextPtr == "" || services_map[*deleteTextPtr] == false {
			fmt.Printf(
				"Vagrant Directory %s or Service %s not found\n", *deleteDircPtr, *deleteTextPtr)
			funcs.PrintInvalid()
		}

		funcs.DeleteONAPComponent(*deleteDircPtr, *deleteTextPtr)
	}
}
