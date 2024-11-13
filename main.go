package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/frafra800/terraform-provider-terraprovider/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,  // Pass the Provider function from the 'provider' package
	})
}