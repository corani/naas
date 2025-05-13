package reasons

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet_ReturnsReason(t *testing.T) {
	reason := Get()

	require.NotEmpty(t, reason,
		"Get() returned an empty string")
}

func TestGet_ReturnsDifferentReasons(t *testing.T) {
	seen := make(map[string]bool)

	for i := 0; i < 100; i++ {
		reason := Get()
		seen[reason] = true
	}

	require.GreaterOrEqual(t, len(seen), 2,
		"Get() did not return different reasons over multiple calls")
}
