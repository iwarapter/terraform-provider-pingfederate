package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/iwarapter/terraform-provider-pingfederate/internal/framework"
)

//go:generate terraform fmt -recursive ./examples/

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	ctx := context.Background()
	//upgradedSdkProvider, err := tf5to6server.UpgradeServer(ctx, sdkv2provider.Provider().GRPCProvider)

	providers := []func() tfprotov6.ProviderServer{
		//func() tfprotov6.ProviderServer {
		//	return upgradedSdkProvider
		//},
		func() tfprotov6.ProviderServer {
			return tfsdk.NewProtocol6Server(framework.New("version")())
		},
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		log.Fatalln(err.Error())
	}

	// Use the result to start a muxed provider
	err = tf6server.Serve("registry.terraform.io/iwarapter/pingfederate", muxServer.ProviderServer)

	if err != nil {
		log.Fatalln(err.Error())
	}
	//
	//var debugMode bool
	//
	//flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	//flag.Parse()
	//
	//if debugMode {
	//	err := plugin.Debug(context.Background(), "registry.terraform.io/-/pingfederate",
	//		&plugin.ServeOpts{
	//			ProviderFunc: pingfederate.Provider,
	//		})
	//	if err != nil {
	//		log.Println(err.Error())
	//	}
	//} else {
	//	plugin.Serve(&plugin.ServeOpts{
	//		ProviderFunc: pingfederate.Provider})
	//}
}
