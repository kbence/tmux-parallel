package tmux

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type TmuxController struct {
	SessionID string
	Cmd       *exec.Cmd
	Channels  struct {
		Stdin  io.WriteCloser
		Stdout *bufio.Reader
	}
}

func NewTmuxController() *TmuxController {
	return &TmuxController{
		SessionID: generateRandomSessionID("tmux-parallel"),
		Cmd:       nil,
	}
}

func (t *TmuxController) NewSession(command string) {
	tmuxPath, err := exec.LookPath("tmux")
	handlePanic(err)

	t.Cmd = exec.Command(tmuxPath, "-C", "new-session", command)

	t.Channels.Stdin, err = t.Cmd.StdinPipe()
	handlePanic(err)

	stdout, err := t.Cmd.StdoutPipe()
	handlePanic(err)
	t.Channels.Stdout = bufio.NewReader(stdout)

	handlePanic(t.Cmd.Start())
	t.readResponse()
}

func (t *TmuxController) Attach() {
	tmuxPath, err := exec.LookPath("tmux")
	handlePanic(err)

	cmd := exec.Command(tmuxPath, "attach", "-t", t.SessionID)
	if err := cmd.Start(); err != nil {
		panic(err)
	}
}

func (t *TmuxController) SplitWindow() {}

func (t *TmuxController) SelectLayout() {}

func (t *TmuxController) NextEvent() {
	fmt.Println(bufio.NewReader(t.Channels.Stdout).ReadLine())
}

func (t *TmuxController) readResponse() {
	for {
		if bytes, _, err := t.Channels.Stdout.ReadLine(); err == nil {
			args := strings.SplitN(string(bytes), " ", 3)
			fmt.Println(args)
		} else {
			break
		}
	}
}

func (t *TmuxController) readEvent() {
	if bytes, _, err := t.Channels.Stdout.ReadLine(); err == nil {
	args := strings.SplitN(string(bytes), " ", 3)
}
