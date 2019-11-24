package main

import (
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
			renderer := NewCommandRenderer(commandTemplate...)

			for _, value := range args[firstValue:] {
				session.RunCommand(renderer.Render(value)...)
			}

			return nil
		},
	}
	cmd.Execute()
}
