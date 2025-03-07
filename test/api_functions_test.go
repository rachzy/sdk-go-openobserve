/*
openobserve

Testing FunctionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/rachzy/sdk-go-openobserve"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_FunctionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FunctionsAPIService CreateFunction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.FunctionsAPI.CreateFunction(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService CreateUpdateEnrichmentTable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var tableName string

		resp, httpRes, err := apiClient.FunctionsAPI.CreateUpdateEnrichmentTable(context.Background(), orgId, tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService DeleteFunction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var name string

		resp, httpRes, err := apiClient.FunctionsAPI.DeleteFunction(context.Background(), orgId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService FunctionPipelineDependency", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var name string

		resp, httpRes, err := apiClient.FunctionsAPI.FunctionPipelineDependency(context.Background(), orgId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService ListFunctions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.FunctionsAPI.ListFunctions(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService TestFunction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.FunctionsAPI.TestFunction(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService UpdateFunction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var name string

		resp, httpRes, err := apiClient.FunctionsAPI.UpdateFunction(context.Background(), orgId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
