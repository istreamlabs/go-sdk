/*
WBD Aventus Channels API

Testing SourcePreviewsAPIService

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

func Test_isp_SourcePreviewsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SourcePreviewsAPIService GetSourcePreviewStream", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var sourceId string

		resp, httpRes, err := apiClient.SourcePreviewsAPI.GetSourcePreviewStream(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcePreviewsAPIService GetSourcePreviewTranscoderStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var sourceId string

		resp, httpRes, err := apiClient.SourcePreviewsAPI.GetSourcePreviewTranscoderStatus(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcePreviewsAPIService PutSourcePreview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var sourceId string

		httpRes, err := apiClient.SourcePreviewsAPI.PutSourcePreview(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcePreviewsAPIService SourcePreviewPinIngest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var sourceId string

		httpRes, err := apiClient.SourcePreviewsAPI.SourcePreviewPinIngest(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcePreviewsAPIService SourcePreviewUnpinIngest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var sourceId string

		httpRes, err := apiClient.SourcePreviewsAPI.SourcePreviewUnpinIngest(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
