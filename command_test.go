package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestCommandRenderer(t *testing.T) {
	TestingT(t)
}

type CommandRendererSuite struct{}

var _ = Suite(&CommandRendererSuite{})

func (s *CommandRendererSuite) Test(c *C) {
	renderer := NewCommandRenderer("echo", "{}")
	cmd := renderer.Render("value")

	c.Assert(len(cmd), Equals, 2)
	c.Assert(cmd[0], Equals, "echo")
	c.Assert(cmd[1], Equals, "value")
}
