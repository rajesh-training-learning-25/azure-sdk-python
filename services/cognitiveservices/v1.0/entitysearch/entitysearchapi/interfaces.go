package entitysearchapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/entitysearch"
)

// EntitiesClientAPI contains the set of methods on the EntitiesClient type.
type EntitiesClientAPI interface {
	Search(ctx context.Context, query string, acceptLanguage string, pragma string, userAgent string, clientID string, clientIP string, location string, countryCode string, market string, responseFilter []entitysearch.AnswerType, responseFormat []entitysearch.ResponseFormat, safeSearch entitysearch.SafeSearch, setLang string) (result entitysearch.SearchResponse, err error)
}

var _ EntitiesClientAPI = (*entitysearch.EntitiesClient)(nil)
