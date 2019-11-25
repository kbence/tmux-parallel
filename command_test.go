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

func (s *CommandRendererSuite) TestSimpleValueRendering(c *C) {
	renderer := NewCommandRenderer("echo", "{}")
	cmd := renderer.Render([]string{"value"})

	c.Assert(cmd, DeepEquals, []string{"echo", "value"})
}

func (s *CommandRendererSuite) TestSimpleOutput(c *C) {
	renderer := NewCommandRenderer("echo")
	cmd := renderer.Render([]string{"value"})

	c.Assert(cmd, DeepEquals, []string{"echo", "value"})
}
