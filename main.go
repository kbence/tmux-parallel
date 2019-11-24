package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:  "tmux-parallel",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			sessionName := fmt.Sprintf("tmux-parallel-%d", rand.Int31())
			sessionExists := false

			tmuxPath, err := exec.LookPath("tmux")

			if err != nil {
				fmt.Println("Cannot find `tmux` on PATH! Exiting")
				os.Exit(1)
			}

			commandTemplate := []string{}
			firstValue := 0

			for idx, arg := range args {
				if arg == ":::" {
					firstValue = idx + 1
					break
				}

				commandTemplate = append(commandTemplate, arg)
			}

			for _, value := range args[firstValue:] {
				command := []string{}
				for _, arg := range commandTemplate {
					command = append(command, strings.ReplaceAll(arg, "{}", value))
				}

				var cmd *exec.Cmd
				if sessionExists {
					cmd = exec.Command(tmuxPath, append([]string{"split-window", "-d", "-t", sessionName}, command...)...)
				} else {
					cmd = exec.Command(tmuxPath, append([]string{"new-session", "-d", "-s", sessionName}, command...)...)
					sessionExists = true
				}

				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Env = os.Environ()
				cmd.Run()
			}

			return nil
		},
	}
	cmd.Execute()
}
