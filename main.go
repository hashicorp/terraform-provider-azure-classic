package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-azure/azure"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: azure.Provider})
}
