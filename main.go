package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/iwarapter/terraform-provider-pingfederate/internal/framework"
	"github.com/iwarapter/terraform-provider-pingfederate/internal/sdkv2provider"
)

//go:generate terraform fmt -recursive ./examples/

var (
	// Version can be updated by goreleaser on release
	version string = "dev"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	flag.Parse()

	ctx := context.Background()
	upgradedSdkProvider, err := tf5to6server.UpgradeServer(ctx, sdkv2provider.Provider().GRPCProvider)
	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedSdkProvider
		},
		providerserver.NewProtocol6(framework.New(version)),
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt

	if *debugFlag {
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	err = tf6server.Serve(
		"registry.terraform.io/iwarapter/pingfederate",
		muxServer.ProviderServer,
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
	//ctx := context.Background()
	////upgradedSdkProvider, err := tf5to6server.UpgradeServer(ctx, sdkv2provider.Provider().GRPCProvider)
	//
	//providers := []func() tfprotov6.ProviderServer{
	//	//func() tfprotov6.ProviderServer {
	//	//	return upgradedSdkProvider
	//	//},
	//	func() tfprotov6.ProviderServer {
	//		return tfsdk.NewProtocol6Server(framework.New("version")())
	//	},
	//}
	//
	//muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)
	//
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//
	//// Use the result to start a muxed provider
	//err = tf6server.Serve("registry.terraform.io/iwarapter/pingfederate", muxServer.ProviderServer)
	//
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
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
