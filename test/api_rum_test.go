/*
openobserve

Testing RumAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openobserve "github.com/rachzy/sdk-go-openobserve"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_RumAPIService(t *testing.T) {

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)

	t.Run("Test RumAPIService LogIngestionJson", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.RumAPI.LogIngestionJson(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RumAPIService ReplayIngestionJson", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.RumAPI.ReplayIngestionJson(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RumAPIService RumIngestionMulti", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.RumAPI.RumIngestionMulti(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
