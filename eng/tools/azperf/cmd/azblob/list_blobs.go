// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/eng/tools/azperf/internal/perf"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/spf13/cobra"
)

var ListBlobCmd = &cobra.Command{
	Use:   "BlobListTest",
	Short: "BlobListTest performance test",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(c *cobra.Command, args []string) error {
		return perf.RunPerfTest(&listBlobPerfTest{})
	},
}

type listBlobPerfTest struct {
	containerName string
	blobName        string
	containerClient azblob.ContainerClient
	blobClient      azblob.BlockBlobClient
}

func (m *listBlobPerfTest) GlobalSetup(ctx context.Context) error {
	m.blobName = "listTest"
	m.containerName = "listtest"

	connStr, ok := os.LookupEnv("AZURE_STORAGE_CONNECTION_STRING")
	if !ok {
		return fmt.Errorf("the environment variable 'AZBLOB_CONNECTION_STRING' could not be found")
	}

	containerClient, err := azblob.NewContainerClientFromConnectionString(connStr, m.containerName, nil)
	if err != nil {
		panic(err)
	}
	m.containerClient = containerClient
	_, err = m.containerClient.Create(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	m.blobClient = m.containerClient.NewBlockBlobClient(m.blobName)

	for i := 0; i < count; i++ {
		_, err = m.blobClient.Upload(
			context.Background(),
			NopCloser(bytes.NewReader([]byte(fmt.Sprintf("listTest%d", i)))),
			nil,
		)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (m *listBlobPerfTest) GlobalTearDown(ctx context.Context) error {
	_, err := m.containerClient.Delete(context.Background(), nil)
	return err
}

func (m *listBlobPerfTest) Setup(ctx context.Context) error {
	return nil
}

func (m *listBlobPerfTest) Run(ctx context.Context) error {
	pager := m.containerClient.ListBlobsFlat(&azblob.ContainerListBlobFlatSegmentOptions{Maxresults: to.Int32Ptr(int32(count))})
	for pager.NextPage(context.Background()) {
		for _ = range pager.PageResponse().ContainerListBlobFlatSegmentResult.Segment.BlobItems {
		}
	}
	return pager.Err()
}

func (m *listBlobPerfTest) TearDown(ctx context.Context) error {
	return nil
}

func (m *listBlobPerfTest) GetMetadata() string {
	return "BlobListTest"
}
