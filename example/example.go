/*
	iStreamPlanet Go SDK Example. Run via:

	CLIENT_ID=... CLIENT_SECRET=... ORG=... go run ./example
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/istreamlabs/go-sdk/isp"
)

func main() {
	client := isp.NewWithClientCredentials(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("ORG"))

	ctx := context.Background()

	fmt.Println("Listing sources:")
	sourceSummaries, _, err := client.SourcesApi.ListSources(ctx).Execute()
	if err.Error() != "" {
		panic(err)
	}

	fmt.Printf("Found %d sources\n", len(sourceSummaries))
	for _, s := range sourceSummaries {
		fmt.Println(s.Id + " " + *s.Self)
	}

	fmt.Println("======================")

	fmt.Println("Listing channels:")
	// List channels with a custom optional argument (page size).
	reqChannels := client.ChannelsApi.ListChannels(ctx).PageSize(2)
	chanSummaries, _, err := reqChannels.Execute()
	if err.Error() != "" {
		panic(err)
	}

	fmt.Printf("Found %d channels\n", len(chanSummaries))
	for _, ch := range chanSummaries {
		fmt.Println(ch.Id + " " + *ch.Self)
	}

	fmt.Println("======================")

	if len(chanSummaries) > 0 {
		fmt.Println("Getting channel details via self link:")

		var ch isp.Channel
		_, err2 := client.GetModel(*chanSummaries[0].Self, &ch)
		if err2 != nil {
			panic(err)
		}

		data, _ := json.MarshalIndent(ch, "", "  ")
		fmt.Println(string(data))
	}
}
