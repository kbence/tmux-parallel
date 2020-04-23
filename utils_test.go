package main

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func TestUtils(t *testing.T) {
	TestingT(t)
}

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TestReadLinesFromFileReturnssLines(c *C) {
	tmpDir := newTempDir()
	defer tmpDir.Release()
	tmpDir.CreateFile("fruits", "apple\norange\npeach\n")

	freader := fileReader{}
	lines := freader.ReadLinesFromFile(tmpDir.PathOf("/fruits"))

	c.Assert(lines, DeepEquals, []string{"apple", "orange", "peach"})
}

func (s *UtilsSuite) TestReadLinesPanicsOnError(c *C) {
	tmpDir := newTempDir()
	defer tmpDir.Release()

	defer func() {
		r := recover()
		c.Assert(r, Not(Equals), nil)
		c.Assert(r.(*os.PathError), Not(Equals), nil)
	}()

	freader := fileReader{}
	freader.ReadLinesFromFile(tmpDir.Path + "/non-existing-file")
}
