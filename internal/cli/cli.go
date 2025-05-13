package cli

import (
	"fmt"
	"io"
)

func Run(w io.Writer, getter func() string) {
	//nolint:errcheck
	fmt.Fprintln(w, getter())
}
