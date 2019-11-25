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

func (s *ArgumentExpanderSuite) TestSimpleExpansion(c *C) {
	args := ParseArgumentExpansion([]string{"a", "b", "c"})

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"a"},
		[]string{"b"},
		[]string{"c"},
	})
}

func (s *ArgumentExpanderSuite) TestMultiDimExpansion(c *C) {
	args := ParseArgumentExpansion([]string{"a", "b", ":::", "c", "d"})

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"a", "c"},
		[]string{"a", "d"},
		[]string{"b", "c"},
		[]string{"b", "d"},
	})
}

func (s *ArgumentExpanderSuite) TestMultiColumnExpansion(c *C) {
	args := ParseArgumentExpansion([]string{"a", "b", ":::+", "c", "d"})

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"a", "c"},
		[]string{"b", "d"},
	})
}

func (s *ArgumentExpanderSuite) TestComplexColumnExpansion(c *C) {
	args := ParseArgumentExpansion(
		[]string{
			"a", "b",
			":::+", "c", "d",
			":::", "e", "f", "g",
			":::+", "h", "i", "j",
		},
	)

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"a", "c", "e", "h"},
		[]string{"a", "c", "f", "i"},
		[]string{"a", "c", "g", "j"},
		[]string{"b", "d", "e", "h"},
		[]string{"b", "d", "f", "i"},
		[]string{"b", "d", "g", "j"},
	})
}

func assertArgumentsExpandedTo(c *C, args *ArgumentExpander, result [][]string) {
	for _, row := range result {
		c.Assert(args.Next(), Equals, true)
		c.Assert(args.Value(), DeepEquals, row)
	}

	c.Assert(args.Next(), Equals, false)
}
