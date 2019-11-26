package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var listPanesRegexp = regexp.MustCompile("^([^:]+)\\:(\\d+)\\.(\\d+)")

// Tmux - struct to manage `tmux` sessions
type Tmux struct {
	BinaryPath string
	SessionID  string

	sessionExists bool
	attachCommand *exec.Cmd
}

type TmuxPane struct {
	SessionID string
	WindowID  int
	PaneID    int
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
		cmd = commandRedirect(t.BinaryPath, "split-window", "-d", "-t", t.SessionID, commandStr)
	} else {
		cmd = commandRedirect(t.BinaryPath, "new-session", "-d", "-s", t.SessionID, commandStr)
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

func (t *Tmux) ListPanes() []*TmuxPane {
	panes := []*TmuxPane{}

	cmd := exec.Command(t.BinaryPath, "list-panes", "-a", "-t", t.SessionID)
	output, err := cmd.Output()

	if err != nil {
		return panes
	}

	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		m := listPanesRegexp.FindStringSubmatch(line)

		if len(m) > 0 {
			panes = append(panes, &TmuxPane{
				SessionID: m[1],
				WindowID:  atoi(m[2]),
			})
		}
	}

	return panes
}

// GetPaneCount - Returns the number of available panes
func (t *Tmux) GetPaneCount() int {
	return len(t.ListPanes())
}

// AttachAsync - Attaches the terminal to tmux and returns (non-blocking)
func (t *Tmux) AttachAsync() {
	t.attachCommand = commandRedirect(t.BinaryPath, "attach", "-t", t.SessionID)
	t.attachCommand.Start()
}

// IsAttached - returns if the process is running (since it means our
// terminal should be attached)
func (t *Tmux) IsAttached() bool {
	if t.attachCommand == nil {
		return false
	}

	_, err := os.FindProcess(t.attachCommand.Process.Pid)

	return err == nil

	// return !t.attachCommand.ProcessState.Exited()
}

// Wait - Waits until tmux exists. Should be called after AttachAsync
func (t *Tmux) Wait() {
	t.attachCommand.Wait()
}
