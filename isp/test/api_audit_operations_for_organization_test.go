/*
WBD Aventus Channels API

Testing AuditOperationsForOrganizationApiService

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

func Test_isp_AuditOperationsForOrganizationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuditOperationsForOrganizationApiService GetOrgChannelTimeline", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var channelId string

		resp, httpRes, err := apiClient.AuditOperationsForOrganizationApi.GetOrgChannelTimeline(context.Background(), org, channelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
