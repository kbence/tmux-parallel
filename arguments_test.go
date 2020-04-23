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

func (s *ArgumentExpanderSuite) TestCartesianExpansion(c *C) {
	args := ParseArgumentExpansion([]string{":::", "a", "b", "c"})

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"a"},
		[]string{"b"},
		[]string{"c"},
	})
}

func (s *ArgumentExpanderSuite) TestMultiDimExpansion(c *C) {
	args := ParseArgumentExpansion([]string{":::", "a", "b", ":::", "c", "d"})

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"a", "c"},
		[]string{"a", "d"},
		[]string{"b", "c"},
		[]string{"b", "d"},
	})
}

func (s *ArgumentExpanderSuite) TestLinkedColumnExpansion(c *C) {
	args := ParseArgumentExpansion([]string{":::", "a", "b", ":::+", "c", "d"})

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"a", "c"},
		[]string{"b", "d"},
	})
}

func (s *ArgumentExpanderSuite) TestComplexColumnExpansion(c *C) {
	args := ParseArgumentExpansion(
		[]string{
			":::", "a", "b",
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

func (s *ArgumentExpanderSuite) TestCartesianFileExpansion(c *C) {
	tmpDir := newTempDir()
	defer tmpDir.Release()

	tmpDir.CreateFile("fruits", "apple\norange\npeach")

	args := ParseArgumentExpansion(
		[]string{
			":::", "a", "b",
			"::::", tmpDir.PathOf("fruits"),
		},
	)

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"a", "apple"},
		[]string{"a", "orange"},
		[]string{"a", "peach"},
		[]string{"b", "apple"},
		[]string{"b", "orange"},
		[]string{"b", "peach"},
	})
}

func (s *ArgumentExpanderSuite) TestLinkedFileExpansion(c *C) {
	tmpDir := newTempDir()
	defer tmpDir.Release()

	tmpDir.CreateFile("drinks", "beer\nwine\ngin")
	tmpDir.CreateFile("fruits", "apple\norange\npeach\nstrawberry")

	args := ParseArgumentExpansion(
		[]string{
			"::::", tmpDir.PathOf("drinks"),
			"::::+", tmpDir.PathOf("fruits"),
		},
	)

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"beer", "apple"},
		[]string{"wine", "orange"},
		[]string{"gin", "peach"},
	})
}

func (s *ArgumentExpanderSuite) TestCartesianMultiFileExpansion(c *C) {
	tmpDir := newTempDir()
	defer tmpDir.Release()

	tmpDir.CreateFile("drinks", "beer\nwine")
	tmpDir.CreateFile("fruits", "apple\norange")

	args := ParseArgumentExpansion(
		[]string{
			"::::", tmpDir.PathOf("drinks"), tmpDir.PathOf("fruits"),
		},
	)

	assertArgumentsExpandedTo(c, args, [][]string{
		[]string{"beer", "apple"},
		[]string{"beer", "orange"},
		[]string{"wine", "apple"},
		[]string{"wine", "orange"},
	})
}

func assertArgumentsExpandedTo(c *C, args *ArgumentExpander, result [][]string) {
	for _, row := range result {
		c.Assert(args.Next(), Equals, true)
		c.Assert(args.Value(), DeepEquals, row)
	}

	c.Assert(args.Next(), Equals, false)
}
