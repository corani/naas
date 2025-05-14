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
	want := "hello, world"

	require.Contains(t, got, want,
		"Run() output should contain the reason text")
}
