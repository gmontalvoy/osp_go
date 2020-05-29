package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gophercloud/gophercloud"

	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
)

func main() {
	// Auth
	authOpts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Provider
	provider, err := openstack.AuthenticatedClient(authOpts)

	// Clients
	// Identity
	opts := gophercloud.EndpointOpts{}
	identityClient, err := openstack.NewIdentityV3(provider, opts)

	// Compute

	listOpts := projects.ListOpts{
		Enabled: gophercloud.Enabled,
	}

	allPages, err := projects.List(identityClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		panic(err)
	}

	for _, project := range allProjects {
		p := strings.Contains(project.Name, "-project")
		if p == true {
			fmt.Printf("Project Name: %v\n", project.Name)
		}
	}

}
