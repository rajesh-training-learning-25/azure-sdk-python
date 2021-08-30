// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type cosmosOffers struct {
	connection *cosmosClientConnection
}

type cosmosOffersResponse struct {
	Offers []ThroughputProperties `json:"Offers"`
}

func (c cosmosOffers) ReadThroughputIfExists(
	ctx context.Context,
	targetRID string,
	requestOptions *ThroughputRequestOptions) (ThroughputResponse, error) {
	// TODO: might want to replace with query iterator once that is in
	operationContext := cosmosOperationContext{
		resourceType:    resourceTypeOffer,
		resourceAddress: "",
	}

	path, err := generatePathForNameBased(resourceTypeOffer, "", true)
	if err != nil {
		return ThroughputResponse{}, err
	}

	azResponse, err := c.connection.sendQueryRequest(
		path,
		ctx,
		fmt.Sprintf(`SELECT * FROM c WHERE c.offerResourceId = '%s'`, targetRID),
		operationContext,
		requestOptions,
		nil)
	if err != nil {
		return ThroughputResponse{}, err
	}

	var theOffers cosmosOffersResponse
	err = azResponse.UnmarshalAsJSON(&theOffers)
	if err != nil {
		return ThroughputResponse{}, err
	}

	if len(theOffers.Offers) == 0 {
		return ThroughputResponse{}, newCosmosErrorWithStatusCode(http.StatusNotFound)
	}

	// Now read the individual offer
	operationContext = cosmosOperationContext{
		resourceType:    resourceTypeOffer,
		resourceAddress: strings.ToLower(theOffers.Offers[0].offerId), //RID auth is lowercase
	}

	path, err = generatePathForNameBased(resourceTypeOffer, theOffers.Offers[0].selfLink, false)
	if err != nil {
		return ThroughputResponse{}, err
	}

	azResponse, err = c.connection.sendGetRequest(
		path,
		ctx,
		operationContext,
		requestOptions,
		nil)
	if err != nil {
		return ThroughputResponse{}, err
	}

	return newThroughputResponse(azResponse)
}
