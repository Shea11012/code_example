package base62_test

import (
	"fmt"
	"testing"

	"shor_url/pkg/base62"
)

func TestBase62(t *testing.T) {
	tests := []struct {
		value uint64
	}{
		{1000},
		{1001},
		{1002},
		{1003},
	}

	for _, test := range tests {
		b2 := base62.Uint64ToBase62(test.value)
		value, err := base62.Base62ToUint(b2)
		if err != nil || value != test.value {
			t.Logf("v1: %d, v2: %d,error: %v\n", value, test.value, err)
		}
	}
}

func TestUint64ToBase62(t *testing.T) {
	fmt.Println(base62.Uint64ToBase62(0))
}
