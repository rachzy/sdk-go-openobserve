/*
openobserve

Testing AlertsAPIService

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

func Test_openapi_AlertsAPIService(t *testing.T) {

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)

	t.Run("Test AlertsAPIService CreateAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.AlertsAPI.CreateAlert(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService CreateDestination", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.AlertsAPI.CreateDestination(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService DeleteAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var alertId string

		resp, httpRes, err := apiClient.AlertsAPI.DeleteAlert(context.Background(), orgId, alertId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService DeleteAlertDestination", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var destinationName string

		resp, httpRes, err := apiClient.AlertsAPI.DeleteAlertDestination(context.Background(), orgId, destinationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService DeleteAlert_1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var streamName string
		var alertName string

		resp, httpRes, err := apiClient.AlertsAPI.DeleteAlert_0(context.Background(), orgId, streamName, alertName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService EnableAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var alertId string

		resp, httpRes, err := apiClient.AlertsAPI.EnableAlert(context.Background(), orgId, alertId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService EnableAlert_2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var streamName string
		var alertName string

		resp, httpRes, err := apiClient.AlertsAPI.EnableAlert_0(context.Background(), orgId, streamName, alertName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService GetAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var alertId string

		resp, httpRes, err := apiClient.AlertsAPI.GetAlert(context.Background(), orgId, alertId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService GetAlert_3", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var streamName string
		var alertName string

		resp, httpRes, err := apiClient.AlertsAPI.GetAlert_0(context.Background(), orgId, streamName, alertName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService GetDestination", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var destinationName string

		resp, httpRes, err := apiClient.AlertsAPI.GetDestination(context.Background(), orgId, destinationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService ListAlerts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.AlertsAPI.ListAlerts(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService ListAlerts_4", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.AlertsAPI.ListAlerts_0(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService ListDestinations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.AlertsAPI.ListDestinations(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService ListStreamAlerts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var streamName string

		resp, httpRes, err := apiClient.AlertsAPI.ListStreamAlerts(context.Background(), orgId, streamName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService MoveAlerts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.AlertsAPI.MoveAlerts(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService SaveAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var streamName string

		resp, httpRes, err := apiClient.AlertsAPI.SaveAlert(context.Background(), orgId, streamName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService TriggerAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var alertId string

		resp, httpRes, err := apiClient.AlertsAPI.TriggerAlert(context.Background(), orgId, alertId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService TriggerAlert_5", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var streamName string
		var alertName string

		resp, httpRes, err := apiClient.AlertsAPI.TriggerAlert_0(context.Background(), orgId, streamName, alertName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService UpdateAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var alertId string

		resp, httpRes, err := apiClient.AlertsAPI.UpdateAlert(context.Background(), orgId, alertId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService UpdateAlert_6", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var streamName string
		var alertName string

		resp, httpRes, err := apiClient.AlertsAPI.UpdateAlert_0(context.Background(), orgId, streamName, alertName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertsAPIService UpdateDestination", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var destinationName string

		resp, httpRes, err := apiClient.AlertsAPI.UpdateDestination(context.Background(), orgId, destinationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
