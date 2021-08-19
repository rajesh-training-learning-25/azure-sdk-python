package common

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest"
	"github.com/Azure/azure-sdk-for-go/tools/internal/sdk"
)

var (
	cache map[string]autorest.GenerationMetadata
)

func ContainsPackage(root, readme, tag string) (autorest.GenerationMetadata, bool) {
	if cache == nil {
		if err := initCache(root); err != nil {
			panic(err)
		}
	}

	v, ok := cache[fmt.Sprintf("%s|%s", readme, tag)]
	return v, ok
}

func initCache(root string) error {
	cache = make(map[string]autorest.GenerationMetadata)
	m, err := autorest.CollectGenerationMetadata(sdk.ServicesPath(root))
	if err != nil {
		return err
	}

	for _, metadata := range m {
		cache[fmt.Sprintf("%s|%s", metadata.RelativeReadme(), metadata.Tag)] = metadata
	}

	return nil
}
