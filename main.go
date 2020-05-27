package main

import (
	"fmt"
	"os"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

func main() {
	// Auth
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		panic(err)
	}

	// Provider
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}

	// Client
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		panic(err)
	}

	// Listing
	listOpts := servers.ListOpts{
		AllTenants: true,
	}

	allPages, err := servers.List(computeClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}

	for _, server := range allServers {
		fmt.Printf("%+v\n", server)
	}
}
