package arg_parser

import (
	"fmt"
	"os"
)


func CreateONAPComponent(directory string, component string)  {
	os.Chdir(directory)
	fmt.Printf("Doing vagrant up %s from directory %s\n", component, directory)
	VagrantUp(component)
}

func CreateONAPcomponent_DockerBuild(directory string, component string)  {
	// TODO do vagrant up <component> and do stuff for DockerBuild
	fmt.Printf("Doing vagrant up with DockerBuild %s\n", component)
}

func CreateONAPComponent_DockerRun(directory string, component string)  {
	// TODO do vagrant up <component> and do stuff for DockerRun
	fmt.Printf("Doing vagrant up with DockerRun %s\n", component)
}

func DeleteONAPComponent(directory string, component string){
	os.Chdir(directory)
	fmt.Printf("Doing vagrant destroy %s from directory %s\n", component, directory)
	VagrantDestroy(component)
}

func VagrantUp(value string){
	cmdName := "vagrant"
	cmdArgs := []string{"up", value}

	fmt.Printf("Running vagrant up %s", value)

	RunShell(cmdName, cmdArgs)
}

func VagrantDestroy(value string)  {
	cmdName := "vagrant"
	cmdArgs := []string{"destroy", value, "-f"}

	fmt.Printf("Running vagrant destroy %s -f", value)

	RunShell(cmdName, cmdArgs)
}