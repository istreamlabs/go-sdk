/*
	iStreamPlanet Go SDK Example. Run via:

	CLIENT_ID=... CLIENT_SECRET=... ORG=... go run ./example
*/

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/istreamlabs/go-sdk/isp"
)

func main() {
	client := isp.NewWithClientCredentials(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("ORG"))

	fmt.Println("Listing sources:")
	reqSources := client.SourcesApi.ListSources(context.Background())
	sourceSummaries, _, err := client.ListSourcesAll(reqSources)
	if err.Error() != "" {
		panic(err)
	}

	for _, s := range sourceSummaries {
		fmt.Println(s.Id + " " + *s.Self)
	}

	fmt.Println("======================")

	fmt.Println("Listing channels:")
	reqChannels := client.ChannelsApi.ListChannels(context.Background())
	chanSummaries, _, err := client.ListChannelsAll(reqChannels)
	if err.Error() != "" {
		panic(err)
	}

	for _, ch := range chanSummaries {
		fmt.Println(ch.Id + " " + *ch.Self)
	}
}
