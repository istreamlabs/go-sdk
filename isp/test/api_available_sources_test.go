/*
WBD Aventus Channels API

Testing AvailableSourcesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package isp

import (
	"context"
	"testing"

	openapiclient "github.com/istreamlabs/go-sdk/isp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_isp_AvailableSourcesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AvailableSourcesApiService GetOrgSource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var sourceId string

		resp, httpRes, err := apiClient.AvailableSourcesApi.GetOrgSource(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvailableSourcesApiService ListOrgSources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string

		resp, httpRes, err := apiClient.AvailableSourcesApi.ListOrgSources(context.Background(), org).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
