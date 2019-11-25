package main

import (
	"fmt"
	"strings"
)

type CommandRenderer struct {
	Template []string
}

func NewCommandRenderer(template ...string) *CommandRenderer {
	return &CommandRenderer{
		Template: template,
	}
}

func (r *CommandRenderer) Render(values []string) []string {
	command := []string{}
	replaced := false

	for _, arg := range r.Template {
		replacedArg := arg

		if strings.Index(arg, "{}") != -1 {
			replacedArg = strings.ReplaceAll(replacedArg, "{}", strings.Join(values, " "))
			replaced = true
		}

		for idx, value := range values {
			mnemonic := fmt.Sprintf("{%d}", idx+1)

			if strings.Index(arg, mnemonic) != -1 {
				replacedArg = strings.ReplaceAll(replacedArg, mnemonic, value)
				replaced = true
			}
		}

		command = append(command, replacedArg)
	}

	if !replaced {
		command = append(command, values...)
	}

	return command
}
