package arg_parser

import "fmt"

func CreateONAPComponent(component string)  {
	// TODO do vagrant up <component>
	fmt.Printf("Doing vagrant up %s\n", component)
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
