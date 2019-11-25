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

func (s *CommandRendererSuite) TestNoReplacement(c *C) {
	renderer := NewCommandRenderer("echo")
	cmd := renderer.Render([]string{"value"})

	c.Assert(cmd, DeepEquals, []string{"echo", "value"})
}

func (s *CommandRendererSuite) TestNumberedReplacement(c *C) {
	renderer := NewCommandRenderer("echo", "{2}", "{1}")
	cmd := renderer.Render([]string{"arg1", "arg2"})

	c.Assert(cmd, DeepEquals, []string{"echo", "arg2", "arg1"})
}
