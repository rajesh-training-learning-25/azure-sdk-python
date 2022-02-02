//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
)

func TestNewPipelineWithOptions(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse()
	opt := arm.ClientOptions{}
	opt.Transport = srv
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	pl := NewPipeline("armtest", "v1.2.3", mockCredential{}, azruntime.PipelineOptions{}, &opt)
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	if ua := resp.Request.Header.Get("User-Agent"); !strings.HasPrefix(ua, "azsdk-go-armtest/v1.2.3") {
		t.Fatalf("unexpected User-Agent %s", ua)
	}
}

func TestNewPipelineWithCustomTelemetry(t *testing.T) {
	const myTelemetry = "something"
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse()
	opt := arm.ClientOptions{}
	opt.Transport = srv
	opt.Telemetry.ApplicationID = myTelemetry
	if opt.Telemetry.ApplicationID != myTelemetry {
		t.Fatalf("telemetry was modified: %s", opt.Telemetry.ApplicationID)
	}
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	pl := NewPipeline("armtest", "v1.2.3", mockCredential{}, azruntime.PipelineOptions{}, &opt)
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	if ua := resp.Request.Header.Get("User-Agent"); !strings.HasPrefix(ua, myTelemetry+" "+"azsdk-go-armtest/v1.2.3") {
		t.Fatalf("unexpected User-Agent %s", ua)
	}
}

func TestDisableAutoRPRegistration(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	// initial response that RP is unregistered
	srv.SetResponse(mock.WithStatusCode(http.StatusConflict), mock.WithBody([]byte(rpUnregisteredResp)))
	opts := &arm.ClientOptions{DisableRPRegistration: true, ClientOptions: policy.ClientOptions{Transport: srv}}
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// log only RP registration
	log.SetEvents(LogRPRegistration)
	defer func() {
		// reset logging
		log.SetEvents()
	}()
	logEntries := 0
	log.SetListener(func(cls log.Event, msg string) {
		logEntries++
	})
	pl := NewPipeline("armtest", "v1.2.3", mockCredential{}, azruntime.PipelineOptions{}, opts)
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusConflict {
		t.Fatalf("unexpected status code %d:", resp.StatusCode)
	}
	// shouldn't be any log entries
	if logEntries != 0 {
		t.Fatalf("expected 0 log entries, got %d", logEntries)
	}
}

// policy that tracks the number of times it was invoked
type countingPolicy struct {
	count int
}

func (p *countingPolicy) Do(req *policy.Request) (*http.Response, error) {
	p.count++
	return req.Next()
}

func TestPipelineWithCustomPolicies(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	// initial response is a failure to trigger retry
	srv.AppendResponse(mock.WithStatusCode(http.StatusInternalServerError))
	srv.AppendResponse(mock.WithStatusCode(http.StatusOK))
	perCallPolicy := countingPolicy{}
	perRetryPolicy := countingPolicy{}
	opts := &arm.ClientOptions{
		DisableRPRegistration: true,
		ClientOptions: policy.ClientOptions{
			PerCallPolicies:  []policy.Policy{&perCallPolicy},
			PerRetryPolicies: []policy.Policy{&perRetryPolicy},
			Retry:            policy.RetryOptions{RetryDelay: time.Microsecond},
			Transport:        srv,
		},
	}
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatal(err)
	}
	pl := NewPipeline("armtest", "v1.2.3", mockCredential{}, azruntime.PipelineOptions{}, opts)
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code %d", resp.StatusCode)
	}
	if perCallPolicy.count != 1 {
		t.Fatalf("unexpected per call policy count %d", perCallPolicy.count)
	}
	if perRetryPolicy.count != 2 {
		t.Fatalf("unexpected per retry policy count %d", perRetryPolicy.count)
	}
}

func TestPipelineAudiences(t *testing.T) {
	for _, c := range []cloud.Name{cloud.AzureChina, cloud.AzureGovernment, cloud.AzurePublicCloud} {
		srv, close := mock.NewServer()
		defer close()
		srv.AppendResponse(mock.WithStatusCode(200))
		opts := &arm.ClientOptions{}
		opts.Cloud = cloud.WellKnownClouds[c]
		opts.Transport = srv
		getTokenCalled := false
		cred := mockCredential{getTokenImpl: func(ctx context.Context, options shared.TokenRequestOptions) (*shared.AccessToken, error) {
			getTokenCalled = true
			if n := len(options.Scopes); n != 1 {
				t.Fatalf("expected 1 scope, got %d", n)
			}
			for _, a := range opts.Cloud.Services[cloud.ResourceManager].Audiences {
				if a == options.Scopes[0] {
					return &shared.AccessToken{Token: "...", ExpiresOn: time.Now().Add(time.Hour)}, nil
				}
			}
			t.Fatalf(`unexpected scope "%s"`, options.Scopes[0])
			return nil, nil
		}}
		req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
		if err != nil {
			t.Fatal(err)
		}
		pl := NewPipeline("test", "v0.1.0", cred, azruntime.PipelineOptions{}, opts)
		_, err = pl.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if !getTokenCalled {
			t.Fatal("mock credential's GetToken method wasn't called")
		}
	}
}

func TestPipelineWithIncompleteCloudConfig(t *testing.T) {
	t.Skip("this test can't pass without a breaking change to arm/runtime.NewPipeline")
	partialConfigs := []cloud.Configuration{
		{Name: "..."},
		{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{"...": {Endpoint: "..."}}},
		{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {Audiences: []string{"..."}},
		}},
		{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {Endpoint: "http://localhost"},
		}},
	}
	for _, c := range partialConfigs {
		opts := &arm.ClientOptions{}
		opts.Cloud = c
		// TODO: uncomment after adding error return to NewPipeline
		// _, err := NewPipeline("test", "v0.1.0", mockCredential{}, azruntime.PipelineOptions{}, opts)
		// if err == nil {
		// 	t.Fatal("expected an error")
		// }
	}
}
