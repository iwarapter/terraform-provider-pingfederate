package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/iwarapter/terraform-provider-pingfederate/pingfederate"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "registry.terraform.io/-/pingfederate",
		ProviderFunc: pingfederate.Provider,
	}

	plugin.Serve(opts)
}
