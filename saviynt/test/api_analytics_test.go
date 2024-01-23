/*
Saviynt Enterprise Identity Cloud (EIC) API

Testing AnalyticsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/saviynt/gosaviynt/saviynt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_AnalyticsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AnalyticsAPIService FetchControlAttributes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AnalyticsAPI.FetchControlAttributes(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsAPIService FetchControlDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AnalyticsAPI.FetchControlDetails(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsAPIService FetchControlDetailsES", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AnalyticsAPI.FetchControlDetailsES(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsAPIService FetchControlList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AnalyticsAPI.FetchControlList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsAPIService FetchControlListES", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AnalyticsAPI.FetchControlListES(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsAPIService FetchRuntimeControlsData", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AnalyticsAPI.FetchRuntimeControlsData(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsAPIService FetchRuntimeControlsDataV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AnalyticsAPI.FetchRuntimeControlsDataV2(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsAPIService RunAnalyticsControls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AnalyticsAPI.RunAnalyticsControls(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
