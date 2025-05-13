package cfg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	// Save original version and restore after test
	orig := version
	t.Cleanup(func() { version = orig })

	version = "1.2.3"
	require.Equal(t, "1.2.3", Version())

	version = "12345/merge"
	require.Equal(t, "pr-12345", Version())

	version = "  2.0.0  "
	require.Equal(t, "2.0.0", Version())
}

func TestHash(t *testing.T) {
	orig := hash
	t.Cleanup(func() { hash = orig })

	hash = "abcdef "
	require.Equal(t, "abcdef", Hash())

	hash = " 12345 "
	require.Equal(t, "12345", Hash())
}

func TestBuild(t *testing.T) {
	orig := build
	t.Cleanup(func() { build = orig })

	build = "20250513 "
	require.Equal(t, "20250513", Build())

	build = " 2025-05-13 10:00:00 "
	require.Equal(t, "2025-05-13 10:00:00", Build())
}
