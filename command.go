package main

import "strings"

type CommandRenderer struct {
	Template []string
}

func NewCommandRenderer(template ...string) *CommandRenderer {
	return &CommandRenderer{
		Template: template,
	}
}

func (r *CommandRenderer) Render(value string) []string {
	command := []string{}

	for _, arg := range r.Template {
		command = append(command, strings.ReplaceAll(arg, "{}", value))
	}

	return command
}