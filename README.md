Vagrant ONAP CLI
----------------

This tool is to provide a simple CLI for the [ONAP Vagrant](http://onap.readthedocs.io/en/latest/submodules/integration.git/bootstrap/vagrant-onap/doc/source/)
project. 

### Example usages:

Help:

`vagrant-onap help`

List of available Vagrant ONAP components:

`vagrant-onap list`

Create component to only clone repos:

`vagrant-onap create -d <VagrantFile path> -component=<name>`

Create component by cloning and building containers:

`vagrant-onap create -d <VagrantFile path> -component=<name> --build`

Create component by cloning, building and running containers:

`vagrant-onap create -d <VagrantFile path> -component=<name> --run`

Delete component:

`vagrant-onap delete -d <VagrantFile path> -component=<name>`
