/*
WBD Aventus Channels API

Testing ChannelOperationsForOrganizationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package isp

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/istreamlabs/go-sdk/isp"
)

func Test_isp_ChannelOperationsForOrganizationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChannelOperationsForOrganizationApiService ClearOrgDvrWindow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var channelId string

		httpRes, err := apiClient.ChannelOperationsForOrganizationApi.ClearOrgDvrWindow(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService GetOrgPreviewImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var channelId string

		httpRes, err := apiClient.ChannelOperationsForOrganizationApi.GetOrgPreviewImage(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService GetOrgSignalLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var channelId string

		resp, httpRes, err := apiClient.ChannelOperationsForOrganizationApi.GetOrgSignalLogs(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService GetOrgSignals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var channelId string

		resp, httpRes, err := apiClient.ChannelOperationsForOrganizationApi.GetOrgSignals(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService InsertOrgId3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var channelId string

		resp, httpRes, err := apiClient.ChannelOperationsForOrganizationApi.InsertOrgId3(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService InsertOrgScte35", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var channelId string

		httpRes, err := apiClient.ChannelOperationsForOrganizationApi.InsertOrgScte35(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService OrgGetTranscoderStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channelId string
		var org string

		resp, httpRes, err := apiClient.ChannelOperationsForOrganizationApi.OrgGetTranscoderStatus(context.Background(), channelId, org).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService OrgPinIngest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channelId string
		var org string

		httpRes, err := apiClient.ChannelOperationsForOrganizationApi.OrgPinIngest(context.Background(), channelId, org).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService OrgPreviewStreams", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var channelId string

		resp, httpRes, err := apiClient.ChannelOperationsForOrganizationApi.OrgPreviewStreams(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService OrgUnpinIngest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channelId string
		var org string

		httpRes, err := apiClient.ChannelOperationsForOrganizationApi.OrgUnpinIngest(context.Background(), channelId, org).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelOperationsForOrganizationApiService PostOrgSignals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var org string
		var channelId string

		resp, httpRes, err := apiClient.ChannelOperationsForOrganizationApi.PostOrgSignals(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
