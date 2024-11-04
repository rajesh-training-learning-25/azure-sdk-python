// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armiotoperations_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
	"log"
)

// Generated from example definition: 2024-11-01/BrokerAuthentication_CreateOrUpdate_Complex.json
func ExampleBrokerAuthenticationClient_BeginCreateOrUpdate_brokerAuthenticationCreateOrUpdateComplex() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBrokerAuthenticationClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", armiotoperations.BrokerAuthenticationResource{
		Properties: &armiotoperations.BrokerAuthenticationProperties{
			AuthenticationMethods: []*armiotoperations.BrokerAuthenticatorMethods{
				{
					Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodServiceAccountToken),
					ServiceAccountTokenSettings: &armiotoperations.BrokerAuthenticatorMethodSat{
						Audiences: []*string{
							to.Ptr("aio-internal"),
						},
					},
				},
				{
					Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodX509),
					X509Settings: &armiotoperations.BrokerAuthenticatorMethodX509{
						TrustedClientCaCert: to.Ptr("my-ca"),
						AuthorizationAttributes: map[string]*armiotoperations.BrokerAuthenticatorMethodX509Attributes{
							"root": {
								Subject: to.Ptr("CN = Contoso Root CA Cert, OU = Engineering, C = US"),
								Attributes: map[string]*string{
									"organization": to.Ptr("contoso"),
								},
							},
							"intermediate": {
								Subject: to.Ptr("CN = Contoso Intermediate CA"),
								Attributes: map[string]*string{
									"city": to.Ptr("seattle"),
									"foo":  to.Ptr("bar"),
								},
							},
							"smart-fan": {
								Subject: to.Ptr("CN = smart-fan"),
								Attributes: map[string]*string{
									"building": to.Ptr("17"),
								},
							},
						},
					},
				},
			},
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.BrokerAuthenticationClientCreateOrUpdateResponse{
	// 	BrokerAuthenticationResource: &armiotoperations.BrokerAuthenticationResource{
	// 		Properties: &armiotoperations.BrokerAuthenticationProperties{
	// 			AuthenticationMethods: []*armiotoperations.BrokerAuthenticatorMethods{
	// 				{
	// 					Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodServiceAccountToken),
	// 					ServiceAccountTokenSettings: &armiotoperations.BrokerAuthenticatorMethodSat{
	// 						Audiences: []*string{
	// 							to.Ptr("aio-internal"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodX509),
	// 					X509Settings: &armiotoperations.BrokerAuthenticatorMethodX509{
	// 						TrustedClientCaCert: to.Ptr("my-ca"),
	// 						AuthorizationAttributes: map[string]*armiotoperations.BrokerAuthenticatorMethodX509Attributes{
	// 							"root": &armiotoperations.BrokerAuthenticatorMethodX509Attributes{
	// 								Subject: to.Ptr("CN = Contoso Root CA Cert, OU = Engineering, C = US"),
	// 								Attributes: map[string]*string{
	// 									"organization": to.Ptr("contoso"),
	// 								},
	// 							},
	// 							"intermediate": &armiotoperations.BrokerAuthenticatorMethodX509Attributes{
	// 								Subject: to.Ptr("CN = Contoso Intermediate CA"),
	// 								Attributes: map[string]*string{
	// 									"city": to.Ptr("seattle"),
	// 									"foo": to.Ptr("bar"),
	// 								},
	// 							},
	// 							"smart-fan": &armiotoperations.BrokerAuthenticatorMethodX509Attributes{
	// 								Subject: to.Ptr("CN = smart-fan"),
	// 								Attributes: map[string]*string{
	// 									"building": to.Ptr("17"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/authentications/resource-name123"),
	// 		Name: to.Ptr("lwucizfvtsdpx"),
	// 		Type: to.Ptr("kvtilkgcxanlfozrd"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-11-01/BrokerAuthentication_CreateOrUpdate_MaximumSet_Gen.json
func ExampleBrokerAuthenticationClient_BeginCreateOrUpdate_brokerAuthenticationCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBrokerAuthenticationClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", armiotoperations.BrokerAuthenticationResource{
		Properties: &armiotoperations.BrokerAuthenticationProperties{
			AuthenticationMethods: []*armiotoperations.BrokerAuthenticatorMethods{
				{
					Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodCustom),
					CustomSettings: &armiotoperations.BrokerAuthenticatorMethodCustom{
						Auth: &armiotoperations.BrokerAuthenticatorCustomAuth{
							X509: &armiotoperations.X509ManualCertificate{
								SecretRef: to.Ptr("secret-name"),
							},
						},
						CaCertConfigMap: to.Ptr("pdecudefqyolvncbus"),
						Endpoint:        to.Ptr("https://www.example.com"),
						Headers: map[string]*string{
							"key8518": to.Ptr("bwityjy"),
						},
					},
					ServiceAccountTokenSettings: &armiotoperations.BrokerAuthenticatorMethodSat{
						Audiences: []*string{
							to.Ptr("jqyhyqatuydg"),
						},
					},
					X509Settings: &armiotoperations.BrokerAuthenticatorMethodX509{
						AuthorizationAttributes: map[string]*armiotoperations.BrokerAuthenticatorMethodX509Attributes{
							"key3384": {
								Attributes: map[string]*string{
									"key186": to.Ptr("ucpajramsz"),
								},
								Subject: to.Ptr("jpgwctfeixitptfgfnqhua"),
							},
						},
						TrustedClientCaCert: to.Ptr("vlctsqddl"),
					},
				},
			},
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.BrokerAuthenticationClientCreateOrUpdateResponse{
	// 	BrokerAuthenticationResource: &armiotoperations.BrokerAuthenticationResource{
	// 		Properties: &armiotoperations.BrokerAuthenticationProperties{
	// 			AuthenticationMethods: []*armiotoperations.BrokerAuthenticatorMethods{
	// 				{
	// 					Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodCustom),
	// 					CustomSettings: &armiotoperations.BrokerAuthenticatorMethodCustom{
	// 						Auth: &armiotoperations.BrokerAuthenticatorCustomAuth{
	// 							X509: &armiotoperations.X509ManualCertificate{
	// 								SecretRef: to.Ptr("secret-name"),
	// 							},
	// 						},
	// 						CaCertConfigMap: to.Ptr("pdecudefqyolvncbus"),
	// 						Endpoint: to.Ptr("https://www.example.com"),
	// 						Headers: map[string]*string{
	// 							"key8518": to.Ptr("bwityjy"),
	// 						},
	// 					},
	// 					ServiceAccountTokenSettings: &armiotoperations.BrokerAuthenticatorMethodSat{
	// 						Audiences: []*string{
	// 							to.Ptr("jqyhyqatuydg"),
	// 						},
	// 					},
	// 					X509Settings: &armiotoperations.BrokerAuthenticatorMethodX509{
	// 						AuthorizationAttributes: map[string]*armiotoperations.BrokerAuthenticatorMethodX509Attributes{
	// 							"key3384": &armiotoperations.BrokerAuthenticatorMethodX509Attributes{
	// 								Attributes: map[string]*string{
	// 									"key186": to.Ptr("ucpajramsz"),
	// 								},
	// 								Subject: to.Ptr("jpgwctfeixitptfgfnqhua"),
	// 							},
	// 						},
	// 						TrustedClientCaCert: to.Ptr("vlctsqddl"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/authentications/resource-name123"),
	// 		Name: to.Ptr("lwucizfvtsdpx"),
	// 		Type: to.Ptr("kvtilkgcxanlfozrd"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-11-01/BrokerAuthentication_Delete_MaximumSet_Gen.json
func ExampleBrokerAuthenticationClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBrokerAuthenticationClient().BeginDelete(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2024-11-01/BrokerAuthentication_Get_MaximumSet_Gen.json
func ExampleBrokerAuthenticationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBrokerAuthenticationClient().Get(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.BrokerAuthenticationClientGetResponse{
	// 	BrokerAuthenticationResource: &armiotoperations.BrokerAuthenticationResource{
	// 		Properties: &armiotoperations.BrokerAuthenticationProperties{
	// 			AuthenticationMethods: []*armiotoperations.BrokerAuthenticatorMethods{
	// 				{
	// 					Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodCustom),
	// 					CustomSettings: &armiotoperations.BrokerAuthenticatorMethodCustom{
	// 						Auth: &armiotoperations.BrokerAuthenticatorCustomAuth{
	// 							X509: &armiotoperations.X509ManualCertificate{
	// 								SecretRef: to.Ptr("secret-name"),
	// 							},
	// 						},
	// 						CaCertConfigMap: to.Ptr("pdecudefqyolvncbus"),
	// 						Endpoint: to.Ptr("https://www.example.com"),
	// 						Headers: map[string]*string{
	// 							"key8518": to.Ptr("bwityjy"),
	// 						},
	// 					},
	// 					ServiceAccountTokenSettings: &armiotoperations.BrokerAuthenticatorMethodSat{
	// 						Audiences: []*string{
	// 							to.Ptr("jqyhyqatuydg"),
	// 						},
	// 					},
	// 					X509Settings: &armiotoperations.BrokerAuthenticatorMethodX509{
	// 						AuthorizationAttributes: map[string]*armiotoperations.BrokerAuthenticatorMethodX509Attributes{
	// 							"key3384": &armiotoperations.BrokerAuthenticatorMethodX509Attributes{
	// 								Attributes: map[string]*string{
	// 									"key186": to.Ptr("ucpajramsz"),
	// 								},
	// 								Subject: to.Ptr("jpgwctfeixitptfgfnqhua"),
	// 							},
	// 						},
	// 						TrustedClientCaCert: to.Ptr("vlctsqddl"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/authentications/resource-name123"),
	// 		Name: to.Ptr("lwucizfvtsdpx"),
	// 		Type: to.Ptr("kvtilkgcxanlfozrd"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-11-01/BrokerAuthentication_ListByResourceGroup_MaximumSet_Gen.json
func ExampleBrokerAuthenticationClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBrokerAuthenticationClient().NewListByResourceGroupPager("rgiotoperations", "resource-name123", "resource-name123", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armiotoperations.BrokerAuthenticationClientListByResourceGroupResponse{
		// 	BrokerAuthenticationResourceListResult: armiotoperations.BrokerAuthenticationResourceListResult{
		// 		Value: []*armiotoperations.BrokerAuthenticationResource{
		// 			{
		// 				Properties: &armiotoperations.BrokerAuthenticationProperties{
		// 					AuthenticationMethods: []*armiotoperations.BrokerAuthenticatorMethods{
		// 						{
		// 							Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodCustom),
		// 							CustomSettings: &armiotoperations.BrokerAuthenticatorMethodCustom{
		// 								Auth: &armiotoperations.BrokerAuthenticatorCustomAuth{
		// 									X509: &armiotoperations.X509ManualCertificate{
		// 										SecretRef: to.Ptr("secret-name"),
		// 									},
		// 								},
		// 								CaCertConfigMap: to.Ptr("pdecudefqyolvncbus"),
		// 								Endpoint: to.Ptr("https://www.example.com"),
		// 								Headers: map[string]*string{
		// 									"key8518": to.Ptr("bwityjy"),
		// 								},
		// 							},
		// 							ServiceAccountTokenSettings: &armiotoperations.BrokerAuthenticatorMethodSat{
		// 								Audiences: []*string{
		// 									to.Ptr("jqyhyqatuydg"),
		// 								},
		// 							},
		// 							X509Settings: &armiotoperations.BrokerAuthenticatorMethodX509{
		// 								AuthorizationAttributes: map[string]*armiotoperations.BrokerAuthenticatorMethodX509Attributes{
		// 									"key3384": &armiotoperations.BrokerAuthenticatorMethodX509Attributes{
		// 										Attributes: map[string]*string{
		// 											"key186": to.Ptr("ucpajramsz"),
		// 										},
		// 										Subject: to.Ptr("jpgwctfeixitptfgfnqhua"),
		// 									},
		// 								},
		// 								TrustedClientCaCert: to.Ptr("vlctsqddl"),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
		// 				},
		// 				ExtendedLocation: &armiotoperations.ExtendedLocation{
		// 					Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
		// 					Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		// 				},
		// 				ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/authentications/resource-name123"),
		// 				Name: to.Ptr("lwucizfvtsdpx"),
		// 				Type: to.Ptr("kvtilkgcxanlfozrd"),
		// 				SystemData: &armiotoperations.SystemData{
		// 					CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
		// 					CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("gnicpuszwd"),
		// 					LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
