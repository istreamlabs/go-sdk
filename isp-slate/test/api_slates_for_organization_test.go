/*
iStreamPlanet Slate Management API

Testing SlatesForOrganizationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package isp

import (
	"context"
	"testing"

	openapiclient "github.com/istreamlabs/go-sdk/isp-slate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_isp_SlatesForOrganizationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SlatesForOrganizationApiService DeleteOrgSlate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var slateId string

		httpRes, err := apiClient.SlatesForOrganizationApi.DeleteOrgSlate(context.Background(), org, slateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlatesForOrganizationApiService GetOrgSlate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var slateId string

		resp, httpRes, err := apiClient.SlatesForOrganizationApi.GetOrgSlate(context.Background(), org, slateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlatesForOrganizationApiService ListOrgSlates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string

		resp, httpRes, err := apiClient.SlatesForOrganizationApi.ListOrgSlates(context.Background(), org).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlatesForOrganizationApiService PutOrgSlate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var org string
		var slateId string

		resp, httpRes, err := apiClient.SlatesForOrganizationApi.PutOrgSlate(context.Background(), org, slateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
