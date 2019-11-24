package main

type CommandLineParser struct {
	CommandTemplate []string
	Arguments       *ArgumentExpander
}

func NewCommandLineParser() *CommandLineParser {
	return &CommandLineParser{}
}

func (p *CommandLineParser) ParseArgs(args []string) {
	var index int
	var arg string

	p.CommandTemplate = []string{}

	for index, arg = range args {
		if arg == ":::" {
			break
		}

		p.CommandTemplate = append(p.CommandTemplate, arg)
	}

	p.Arguments = ParseArgumentExpansion(args[index+1:])
}
