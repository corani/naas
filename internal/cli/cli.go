package cli

import (
	"fmt"
)

func Run(getter func() string) {
	fmt.Println(getter())
}
