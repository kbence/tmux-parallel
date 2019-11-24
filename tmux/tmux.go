package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Tmux - struct to manage `tmux` sessions
type Tmux struct {
	BinaryPath string
	SessionID  string

	sessionExists bool
}

// New - Creates and configures a new Tmux struct.
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

// ExecCommand - Runs a command in a new pane. If a session does not yet
// exist, it creates one.
func (t *Tmux) ExecCommand(command ...string) {
	var cmd *exec.Cmd

	commandStr := strings.Join(command, " ")

	if t.sessionExists {
		cmd = exec.Command(t.BinaryPath, "split-window", "-d", "-t", t.SessionID, commandStr)
	} else {
		cmd = exec.Command(t.BinaryPath, "new-session", "-d", "-s", t.SessionID, commandStr)
		t.sessionExists = true
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Run()

	t.SelectLayout("tiled")
}

// SelectLayout - Sends a select-layout command to the tmux session
func (t *Tmux) SelectLayout(layout string) {
	exec.Command(t.BinaryPath, "select-layout", "-t", t.SessionID, "tiled").Run()
}

func (t *Tmux) Attach() {
	cmd := exec.Command(t.BinaryPath, "attach", "-t", t.SessionID)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Run()
}
