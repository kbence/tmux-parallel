package main

import (
	"strings"

	"github.com/kbence/tmux-parallel/tmux"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:  "tmux-parallel",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			commandTemplate := []string{}
			firstValue := 0

			for idx, arg := range args {
				if arg == ":::" {
					firstValue = idx + 1
					break
				}

				commandTemplate = append(commandTemplate, arg)
			}

			session := tmux.New()

			for _, value := range args[firstValue:] {
				command := []string{}
				for _, arg := range commandTemplate {
					command = append(command, strings.ReplaceAll(arg, "{}", value))
				}

				session.RunCommand(command...)
			}

			return nil
		},
	}
	cmd.Execute()
}
