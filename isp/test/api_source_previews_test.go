/*
iStreamPlanet Channels API

Testing SourcePreviewsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package isp

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/isp"
)

func Test_isp_SourcePreviewsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SourcePreviewsApiService GetSourcePreviewStream", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var sourceId string

		resp, httpRes, err := apiClient.SourcePreviewsApi.GetSourcePreviewStream(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcePreviewsApiService GetSourcePreviewTranscoderStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var sourceId string

		resp, httpRes, err := apiClient.SourcePreviewsApi.GetSourcePreviewTranscoderStatus(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcePreviewsApiService PutSourcePreview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var sourceId string

		httpRes, err := apiClient.SourcePreviewsApi.PutSourcePreview(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcePreviewsApiService SourcePreviewPinIngest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var sourceId string

		httpRes, err := apiClient.SourcePreviewsApi.SourcePreviewPinIngest(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcePreviewsApiService SourcePreviewUnpinIngest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var sourceId string

		httpRes, err := apiClient.SourcePreviewsApi.SourcePreviewUnpinIngest(context.Background(), org, sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}