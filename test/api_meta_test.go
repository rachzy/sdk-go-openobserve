/*
openobserve

Testing MetaAPIService

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

func Test_openapi_MetaAPIService(t *testing.T) {

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)

	t.Run("Test MetaAPIService Healthz", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MetaAPI.Healthz(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
