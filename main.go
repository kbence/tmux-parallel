package main

import (
	"runtime"
	"time"

	"github.com/kbence/tmux-parallel/tmux"
	"github.com/spf13/cobra"
)

func main() {
	var jobCount int

	cmd := &cobra.Command{
		Use:  "tmux-parallel",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			parser := NewCommandLineParser()
			parser.ParseArgs(args)

			session := tmux.New()
			renderer := NewCommandRenderer(parser.CommandTemplate...)

			tick := time.NewTicker(200 * time.Microsecond)
			defer tick.Stop()

			for parser.Arguments.Next() {
				session.ExecCommand(renderer.Render(parser.Arguments.Value())...)

				if !session.IsAttached() {
					session.AttachAsync()
				}

				for session.GetPaneCount() >= jobCount {
					<-tick.C
				}
			}

			session.Wait()

			return nil
		},
	}

	cmd.Flags().SetInterspersed(false)
	cmd.Flags().IntVarP(&jobCount, "jobs", "j", runtime.NumCPU(), "Number of jobs to run at once")
	cmd.Execute()
}
