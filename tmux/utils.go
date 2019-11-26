package tmux

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
)

func generateRandomSessionId(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, rand.Int31())
}

func commandRedirect(args ...string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd
}

func atoi(s string) int {
	if num, err := strconv.Atoi(s); err == nil {
		return num
	}

	return 0
}
