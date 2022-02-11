// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package main

import (
	"github.com/Azure/azure-sdk-for-go/sdk/internal/perf"
)

func main() {
	tests := map[string]perf.MapInterface{
		"UploadBlobTest":   {Register: uploadTestRegister, New: NewUploadTest},
		"ListBlobTest":     {Register: listTestRegister, New: NewListTest},
		"DownloadBlobTest": {Register: downloadTestRegister, New: NewDownloadTest},
	}

	perf.Run(tests)
}
