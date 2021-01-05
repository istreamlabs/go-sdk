package sdk

/*
Package sdk is the official iStreamPlanet SDK for the Go programming language.

Getting Started

The best way to get started working with the SDK is to use `go get` to add the
SDK:

	go get github.com/istreamlabs/go-sdk/isp

Hello iStreamPlanet

This example shows how you can use the SDK to make an API request to list the
sources available to you.

	package main

	import (
		"context"
		"fmt"

		"github.com/istreamlabs/go-sdk/isp"
	)

	func main() {
		// Create a new API client.
		client := isp.NewWithClientCredentials(CLIENT_ID, CLIENT_SECRET, ORGANIZATION)

		// Get a list of all source summaries. Pagination is handled via the
		// `ListSourcesAll()` utility. For calls without pagination you can simply
		// call the `Execute()` method on the request.
		request := client.SourcesApi.ListSources(context.Background())
		summaries, _, err := client.ListSourcesAll(request)
		if err.Error() != "" {
			panic(err)
		}

		// For each source, print out its ID and self link.
		for _, s := range summaries {
			fmt.Println(s.Id + " " + *s.Self)
		}
	}

*/
