/*
openobserve

Testing UsersAPIService

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

func Test_openapi_UsersAPIService(t *testing.T) {

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)

	t.Run("Test UsersAPIService AddUserToOrg", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var emailId string

		resp, httpRes, err := apiClient.UsersAPI.AddUserToOrg(context.Background(), orgId, emailId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService RemoveUserFromOrg", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var emailId string

		resp, httpRes, err := apiClient.UsersAPI.RemoveUserFromOrg(context.Background(), orgId, emailId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UserList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.UsersAPI.UserList(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UserSave", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.UsersAPI.UserSave(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UserUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var emailId string

		resp, httpRes, err := apiClient.UsersAPI.UserUpdate(context.Background(), orgId, emailId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
