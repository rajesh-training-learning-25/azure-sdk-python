package search

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/azure/go-autorest/autorest"
)

const (
	ApiVersion     = "2015-02-28"
	DefaultBaseUri = "https://management.azure.com"
)

// Client that can be used to manage Azure Search services and API keys.
type SearchManagementClient struct {
	autorest.Client
	BaseUri        string
	SubscriptionId string
}

func New(subscriptionId string) SearchManagementClient {
	return NewWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewWithBaseUri(baseUri string, subscriptionId string) SearchManagementClient {
	return SearchManagementClient{
		Client:         autorest.DefaultClient,
		BaseUri:        baseUri,
		SubscriptionId: subscriptionId,
	}
}
