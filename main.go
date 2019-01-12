package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/iwarapter/terraform-provider-pingfederate/pingfederate"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingfederate.Provider})
}
