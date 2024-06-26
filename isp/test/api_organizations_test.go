/*
WBD Aventus Channels API

Testing OrganizationsApiService

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

func Test_isp_OrganizationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrganizationsApiService ListOrgs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrganizationsApi.ListOrgs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
