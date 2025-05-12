package cli

import (
	"fmt"

	"github.wdf.sap.corp/I331555/naasgo/internal/reasons"
)

func Run() {
	fmt.Println(reasons.Get())
}
