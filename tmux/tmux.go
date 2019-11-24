package tmux

import (
	"fmt"
	"os"
	"os/exec"
)

type Tmux struct {
	BinaryPath string
	SessionID  string

	sessionExists bool
}

func New() *Tmux {
	tmuxPath, err := exec.LookPath("tmux")

	if err != nil {
		fmt.Println("Cannot find `tmux` on PATH! Exiting")
		os.Exit(1)
	}

	return &Tmux{
		BinaryPath: tmuxPath,
		SessionID:  generateRandomSessionId("tmux-parallel"),

		sessionExists: false,
	}
}

func (t *Tmux) RunCommand(command ...string) {
	var cmd *exec.Cmd

	if t.sessionExists {
		cmd = exec.Command(t.BinaryPath, append([]string{"split-window", "-d", "-t", t.SessionID}, command...)...)
	} else {
		cmd = exec.Command(t.BinaryPath, append([]string{"new-session", "-d", "-s", t.SessionID}, command...)...)
		t.sessionExists = true
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Run()
}
