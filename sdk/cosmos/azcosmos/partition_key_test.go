// +build !emulator
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"testing"
)

func TestInvalidPartitionKeyValues(t *testing.T) {
	invalidTypes := []interface{}{
		complex64(0),
		complex128(0),
		[]byte(nil),
		[]byte{},
		// whatever other type of struct
		cosmosOperationContext{},
	}

	for _, invalidType := range invalidTypes {
		_, err := NewPartitionKey(invalidType)
		if err == nil {
			t.Errorf("Expected error for partition key type %v", invalidType)
		}
	}
}

func TestValidPartitionKeyValues(t *testing.T) {
	validTypes := map[interface{}]string{
		nil:           "[null]",
		true:          "[true]",
		false:         "[false]",
		"some string": "[\"some string\"]",
		int(10):       "[10]",
		int8(10):      "[10]",
		int16(10):     "[10]",
		int32(10):     "[10]",
		int64(10):     "[10]",
		uint(10):      "[10]",
		uint8(10):     "[10]",
		uint16(10):    "[10]",
		uint32(10):    "[10]",
		uint64(10):    "[10]",
		float32(10.5): "[10.5]",
		float64(10.5): "[10.5]",
	}

	for validType, expectedSerialization := range validTypes {
		pk, err := NewPartitionKey(validType)
		if err != nil {
			t.Errorf("Expected success for partition key type %v and got %v", validType, err)
		}

		if len(pk.partitionKeyInternal.components) != 1 {
			t.Errorf("Expected partition key to have 1 component, but it has %v", len(pk.partitionKeyInternal.components))
		}

		serialization, err := pk.toJsonString()
		if err != nil {
			t.Errorf("Failed to serialize PK for %v, got %v", validType, err)
		}

		if serialization != expectedSerialization {
			t.Errorf("Expected serialization %v, but got %v", expectedSerialization, serialization)
		}
	}
}
