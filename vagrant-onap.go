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
	createTextPtr := createCommand.String("component", "", "Component name. (Required)")
	createBuildContainerPtr := createCommand.Bool("build", false, "Option to build container")
	createRunContainerPtr := createCommand.Bool("run", false, "Option to run container")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
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
		if *createTextPtr == "" || services_map[*createTextPtr] == false {
			fmt.Printf("Service %s not found\n", *createTextPtr)
			funcs.PrintInvalid()
		}

		if *createBuildContainerPtr == false &&
			*createRunContainerPtr == false {
			funcs.CreateONAPComponent(*createTextPtr)
		}

		if *createBuildContainerPtr == true {
			funcs.CreateONAPcomponent_DockerBuild(*createTextPtr)
		}

		if *createRunContainerPtr == true {
			funcs.CreateONAPComponent_DockerRun(*createTextPtr)
		}

	}

	if deleteCommand.Parsed() {
		if *deleteTextPtr == "" || services_map[*deleteTextPtr] == false {
			fmt.Printf("Service %s not found\n", *deleteTextPtr)
			funcs.PrintInvalid()
		}

		funcs.DeleteONAPComponent(*deleteTextPtr)
	}
}
