// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
	"github.com/stretchr/testify/assert"
)

type countPolicy struct {
	callCount int
}

func (p *countPolicy) Do(req *policy.Request) (*http.Response, error) {
	p.callCount += 1
	return req.Next()
}

func TestGlobalEndpointManagerGetWriteEndpoints(t *testing.T) {
	srv, close := mock.NewTLSServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))

	pl := azruntime.NewPipeline("azcosmostest", "v1.0.0", azruntime.PipelineOptions{}, &policy.ClientOptions{Transport: srv})

	client := &Client{endpoint: srv.URL(), pipeline: pl}

	preferredRegions := []string{"West US", "Central US"}

	gem, err := newGlobalEndpointManager(client, preferredRegions, 5*time.Minute)
	assert.NoError(t, err)

	writeEndpoints, err := gem.GetWriteEndpoints()
	assert.NoError(t, err)

	serverEndpoint, err := url.Parse(srv.URL())
	assert.NoError(t, err)

	expectedWriteEndpoints := []url.URL{
		*serverEndpoint,
	}
	assert.Equal(t, expectedWriteEndpoints, writeEndpoints)
}

func TestGlobalEndpointManagerGetReadEndpoints(t *testing.T) {
	srv, close := mock.NewTLSServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))

	pl := azruntime.NewPipeline("azcosmostest", "v1.0.0", azruntime.PipelineOptions{}, &policy.ClientOptions{Transport: srv})

	client := &Client{endpoint: srv.URL(), pipeline: pl}

	preferredRegions := []string{"West US", "Central US"}

	gem, err := newGlobalEndpointManager(client, preferredRegions, 5*time.Minute)
	assert.NoError(t, err)

	readEndpoints, err := gem.GetReadEndpoints()
	assert.NoError(t, err)

	serverEndpoint, err := url.Parse(srv.URL())
	assert.NoError(t, err)

	expectedReadEndpoints := []url.URL{
		*serverEndpoint,
	}
	assert.Equal(t, expectedReadEndpoints, readEndpoints)
}

func TestGlobalEndpointManagerMarkEndpointUnavailableForRead(t *testing.T) {
	srv, close := mock.NewTLSServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))

	pl := azruntime.NewPipeline("azcosmostest", "v1.0.0", azruntime.PipelineOptions{}, &policy.ClientOptions{Transport: srv})

	client := &Client{endpoint: srv.URL(), pipeline: pl}

	preferredRegions := []string{"West US", "Central US"}

	gem, err := newGlobalEndpointManager(client, preferredRegions, 5*time.Minute)
	assert.NoError(t, err)

	endpoint, err := url.Parse(client.endpoint)
	assert.NoError(t, err)

	readEndpoints, err := gem.GetReadEndpoints()
	assert.NoError(t, err)
	print(readEndpoints)

	err = gem.MarkEndpointUnavailableForRead(*endpoint)
	assert.NoError(t, err)

	unavailable := gem.IsEndpointUnavailable(*endpoint, 1)
	assert.True(t, unavailable)
}

func TestGlobalEndpointManagerMarkEndpointUnavailableForWrite(t *testing.T) {
	srv, close := mock.NewTLSServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))

	pl := azruntime.NewPipeline("azcosmostest", "v1.0.0", azruntime.PipelineOptions{}, &policy.ClientOptions{Transport: srv})

	client := &Client{endpoint: srv.URL(), pipeline: pl}

	preferredRegions := []string{"West US", "Central US"}

	gem, err := newGlobalEndpointManager(client, preferredRegions, 5*time.Minute)
	assert.NoError(t, err)

	endpoint, err := url.Parse(client.endpoint)
	assert.NoError(t, err)

	err = gem.MarkEndpointUnavailableForWrite(*endpoint)
	assert.NoError(t, err)

	unavailable := gem.IsEndpointUnavailable(*endpoint, 2)
	assert.True(t, unavailable)
}

func TestGlobalEndpointManagerGetEndpointLocation(t *testing.T) {
	srv, close := mock.NewTLSServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))

	westRegion := accountRegion{
		Name:     "West US",
		Endpoint: srv.URL(),
	}

	properties := accountProperties{
		ReadRegions:                  []accountRegion{westRegion},
		WriteRegions:                 []accountRegion{westRegion},
		EnableMultipleWriteLocations: false,
	}

	jsonString, err := json.Marshal(properties)
	if err != nil {
		t.Fatal(err)
	}
	srv.SetResponse(mock.WithBody(jsonString))

	pl := azruntime.NewPipeline("azcosmostest", "v1.0.0", azruntime.PipelineOptions{}, &policy.ClientOptions{Transport: srv})

	client := &Client{endpoint: srv.URL(), pipeline: pl}

	gem, err := newGlobalEndpointManager(client, []string{}, 5*time.Minute)
	assert.NoError(t, err)

	serverEndpoint, err := url.Parse(srv.URL())
	assert.NoError(t, err)

	err = gem.Update(context.Background())
	assert.NoError(t, err)

	location := gem.GetEndpointLocation(*serverEndpoint)

	expectedLocation := "West US"
	assert.Equal(t, expectedLocation, location)
}

func TestGlobalEndpointManagerGetAccountProperties(t *testing.T) {
	srv, close := mock.NewTLSServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))

	pl := azruntime.NewPipeline("azcosmostest", "v1.0.0", azruntime.PipelineOptions{}, &policy.ClientOptions{Transport: srv})

	client := &Client{endpoint: srv.URL(), pipeline: pl}

	preferredRegions := []string{"West US", "Central US"}

	gem, err := newGlobalEndpointManager(client, preferredRegions, 5*time.Minute)
	assert.NoError(t, err)

	accountProps, err := gem.GetAccountProperties(context.Background())
	assert.NoError(t, err)

	expectedAccountProps := accountProperties{
		ReadRegions:                  nil,
		WriteRegions:                 nil,
		EnableMultipleWriteLocations: false,
	}
	assert.Equal(t, expectedAccountProps, accountProps)
}

func TestGlobalEndpointManagerCanUseMultipleWriteLocations(t *testing.T) {
	srv, close := mock.NewTLSServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))

	pl := azruntime.NewPipeline("azcosmostest", "v1.0.0", azruntime.PipelineOptions{}, &policy.ClientOptions{Transport: srv})

	client := &Client{endpoint: srv.URL(), pipeline: pl}

	preferredRegions := []string{"West US", "Central US"}

	serverEndpoint, err := url.Parse(srv.URL())
	assert.NoError(t, err)

	mockLc := newLocationCache(preferredRegions, *serverEndpoint)
	mockLc.enableMultipleWriteLocations = true
	mockLc.useMultipleWriteLocations = true

	mockGem := globalEndpointManager{
		client:              client,
		preferredLocations:  preferredRegions,
		locationCache:       mockLc,
		refreshTimeInterval: 5 * time.Minute,
	}

	gem, err := newGlobalEndpointManager(client, preferredRegions, 5*time.Minute)
	assert.NoError(t, err)

	// Multiple locations should be false for default GEM
	canUseMultipleWriteLocs := gem.CanUseMultipleWriteLocations()
	assert.False(t, canUseMultipleWriteLocs)

	// Mock GEM with multiple write locations available should show true
	canUseMultipleWriteLocs = mockGem.CanUseMultipleWriteLocations()
	assert.True(t, canUseMultipleWriteLocs)
}

func TestGlobalEndpointManagerConcurrentUpdate(t *testing.T) {
	countPolicy := &countPolicy{}
	srv, closeFunc := mock.NewTLSServer()
	defer closeFunc()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))

	westRegion := accountRegion{
		Name:     "West US",
		Endpoint: srv.URL(),
	}

	properties := accountProperties{
		ReadRegions:                  []accountRegion{westRegion},
		WriteRegions:                 []accountRegion{westRegion},
		EnableMultipleWriteLocations: false,
	}

	jsonString, err := json.Marshal(properties)
	if err != nil {
		t.Fatal(err)
	}
	srv.SetResponse(mock.WithBody(jsonString))

	pl := azruntime.NewPipeline("azcosmostest", "v1.0.0", azruntime.PipelineOptions{PerCall: []policy.Policy{countPolicy}}, &policy.ClientOptions{Transport: srv})
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	assert.NoError(t, err)

	req.SetOperationValue(pipelineRequestOptions{
		isWriteOperation: true,
	})

	client := &Client{endpoint: srv.URL(), pipeline: pl}
	fmt.Println(client.endpoint)

	gem, err := newGlobalEndpointManager(client, []string{}, 5*time.Second)
	assert.NoError(t, err)

	// Call update concurrently and see how many times the policy gets called
	concurrency := 5
	wg := &sync.WaitGroup{}
	wg.Add(concurrency)

	finishCh := make(chan struct{})

	for i := 0; i < concurrency; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// Call the function in each goroutine
			err := gem.Update(context.Background())
			if err != nil {
				print(err)
			}
		}(wg)
	}

	wg.Wait()
	close(finishCh)

	// Set a timeout for the test
	select {
	case <-finishCh:
	// All goroutines have finished
	case <-time.After(2 * time.Minute):
		t.Fatal("Test timed out noob")
	}

	// Check that the function was called only once
	callCount := countPolicy.callCount
	assert.Equal(t, callCount, 1)
}
