package arg_parser

import (
	"fmt"
	"os"
)

func CreateONAPComponent(directory string, component string)  {
	os.Chdir(directory)
	fmt.Printf("Doing vagrant up %s from directory %s\n", component, directory)
	os.Setenv("SKIP_GET_IMAGES", "True")
	os.Setenv("SKIP_INSTALL", "True")
	VagrantUp(component)
}

func CreateONAPcomponent_DockerBuild(directory string, component string)  {
	os.Chdir(directory)
	fmt.Printf("Doing vagrant up with DockerBuild %s\n", component)
	os.Setenv("SKIP_INSTALL", "True")
	VagrantUp(component)
}

func CreateONAPComponent_DockerRun(directory string, component string)  {
	fmt.Printf("Doing vagrant up with DockerRun %s\n", component)
	VagrantUp(component)
}

func DeleteONAPComponent(directory string, component string){
	os.Chdir(directory)
	fmt.Printf("Doing vagrant destroy %s from directory %s\n", component, directory)
	VagrantDestroy(component)
}

func VagrantUp(value string){
	cmdName := "vagrant"
	cmdArgs := []string{"up", value, "--machine-readable"}

	fmt.Printf("Running vagrant up %s", value)

	err := RunShell(cmdName, cmdArgs)
	if err!=nil {
		fmt.Printf("Error in running vagrant up %s. Error: %s", value, err)
	}
}

func VagrantDestroy(value string)  {
	cmdName := "vagrant"
	cmdArgs := []string{"destroy", value, "-f", "--machine-readable"}

	fmt.Printf("Running vagrant destroy %s -f", value)

	err := RunShell(cmdName, cmdArgs)
	if err!=nil {
		fmt.Printf("Error in running vagrant destroy %s. Error: %s", value, err)
	}
}