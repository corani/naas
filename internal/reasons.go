package internal

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"sync"
)

//go:embed reasons.json
var rawReasons []byte

var reasons = sync.OnceValue(func() []string {
	var reasons []string

	if err := json.Unmarshal(rawReasons, &reasons); err != nil {
		panic(err)
	}

	return reasons
})

func GetReason() string {
	reasons := reasons()

	// return a random reason
	return reasons[rand.Intn(len(reasons))]
}
