/*
Render Public API

Testing JobsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package render

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_render_JobsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test JobsApiService GetJob", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string
        var jobId string

        resp, httpRes, err := apiClient.JobsApi.GetJob(context.Background(), serviceId, jobId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test JobsApiService ListJob", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.JobsApi.ListJob(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test JobsApiService PostJob", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.JobsApi.PostJob(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}