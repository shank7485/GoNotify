package arg_parser

import (
	"fmt"
	"os/exec"
	"log"
)

func CreateONAPComponent(component string)  {
	// TODO: Need to be run from vagrant-onap directory to find VagrantFile
	cmd := exec.Command("vagrant", "up", component)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func CreateONAPcomponent_DockerBuild(component string)  {
	// TODO do vagrant up <component> and do stuff for DockerBuild
	fmt.Printf("Doing vagrant up with DockerBuild %s\n", component)
}

func CreateONAPComponent_DockerRun(component string)  {
	// TODO do vagrant up <component> and do stuff for DockerRun
	fmt.Printf("Doing vagrant up with DockerRun %s\n", component)
}

func DeleteONAPComponent(component string){
	// TODO do vagrant destroy <component>
	fmt.Printf("Doing vagrant delete %s\n", component)
}
