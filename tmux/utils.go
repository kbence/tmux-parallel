package tmux

import (
	"fmt"
	"math/rand"
)

func generateRandomSessionId(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, rand.Int31())
}
