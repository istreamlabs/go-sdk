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

		// Create an empty context. You may already have a context from an incoming
		// request and can use that instead.
		ctx := context.Background()

		// Get a list of all source summaries. Pagination is handled
		// automatically and all source summaries are returned in one list.
		summaries, _, err := client.SourcesApi.ListSources(ctx).Execute()
		if err.Error() != "" {
			panic(err)
		}

		// For each source, print out its ID and self link.
		for _, s := range summaries {
			fmt.Println(s.Id + " " + *s.Self)
		}
	}

*/
