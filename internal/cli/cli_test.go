package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun_PrintsGetterOutput(t *testing.T) {
	var buf bytes.Buffer

	getter := func() string { return "hello, world" }

	Run(&buf, getter)

	got := buf.String()
	want := "hello, world\n"

	require.Equal(t, want, got,
		"Run() output mismatch")
}
