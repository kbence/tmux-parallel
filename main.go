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
			parser := NewCommandLineParser()
			parser.ParseArgs(args)

			session := tmux.New()
			renderer := NewCommandRenderer(parser.CommandTemplate...)

			for parser.Arguments.Next() {
				session.ExecCommand(renderer.Render(parser.Arguments.Value())...)
			}

			session.Attach()

			return nil
		},
	}
	cmd.Execute()
}
