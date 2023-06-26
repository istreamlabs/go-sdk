/*
iStreamPlanet Channels API

Testing DeprecatedLive2VODApiService

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

func Test_isp_DeprecatedLive2VODApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeprecatedLive2VODApiService DeprecatedClipGetMp4", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string
		var vodId string
		var clipId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.DeprecatedClipGetMp4(context.Background(), customerId, productId, programId, vodId, clipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedLive2VODApiService DeprecatedClipGetPresentations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string
		var vodId string
		var clipId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.DeprecatedClipGetPresentations(context.Background(), customerId, productId, programId, vodId, clipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedLive2VODApiService DeprecatedClipGetProgramTime", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string
		var vodId string
		var clipId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.DeprecatedClipGetProgramTime(context.Background(), customerId, productId, programId, vodId, clipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedLive2VODApiService DeprecatedClipMakeMp4", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string
		var vodId string
		var clipId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.DeprecatedClipMakeMp4(context.Background(), customerId, productId, programId, vodId, clipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedLive2VODApiService DeprecatedGetClipManifest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string
		var vodId string
		var clipId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.DeprecatedGetClipManifest(context.Background(), customerId, productId, programId, vodId, clipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedLive2VODApiService GetDeprecatedClip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string
		var vodId string
		var clipId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.GetDeprecatedClip(context.Background(), customerId, productId, programId, vodId, clipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedLive2VODApiService GetDeprecatedProgram", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.GetDeprecatedProgram(context.Background(), customerId, productId, programId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedLive2VODApiService ListDeprecatedClips", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string
		var vodId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.ListDeprecatedClips(context.Background(), customerId, productId, programId, vodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedLive2VODApiService ListDeprecatedVods", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var productId string
		var programId string

		resp, httpRes, err := apiClient.DeprecatedLive2VODApi.ListDeprecatedVods(context.Background(), customerId, productId, programId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
