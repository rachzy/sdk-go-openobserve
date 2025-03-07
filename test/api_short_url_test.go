/*
openobserve

Testing ShortUrlAPIService

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

func Test_openapi_ShortUrlAPIService(t *testing.T) {

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)

	t.Run("Test ShortUrlAPIService Retrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shortId string
		var orgId string

		httpRes, err := apiClient.ShortUrlAPI.Retrieve(context.Background(), shortId, orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShortUrlAPIService Shorten", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ShortUrlAPI.Shorten(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
