module github.com/Azure/azure-sdk-for-go/sdk/data/aztables/testdata/perf

go 1.18

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.16.0
	github.com/Azure/azure-sdk-for-go/sdk/data/aztables v1.3.0
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0
)

require (
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace github.com/Azure/azure-sdk-for-go/sdk/data/aztables => ../..
