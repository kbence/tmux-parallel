package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestArgumentExpander(t *testing.T) {
	TestingT(t)
}

type ArgumentExpanderSuite struct{}

var _ = Suite(&ArgumentExpanderSuite{})

func (s *ArgumentExpanderSuite) Test(c *C) {
	args := ParseArgumentExpansion([]string{"a", "b", "c"})

	c.Assert(args.Next(), Equals, true)
	c.Assert(args.Value(), Equals, "a")

	c.Assert(args.Next(), Equals, true)
	c.Assert(args.Value(), Equals, "b")

	c.Assert(args.Next(), Equals, true)
	c.Assert(args.Value(), Equals, "c")

	c.Assert(args.Next(), Equals, false)
}
