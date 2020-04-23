package main

const (
	cartesianExpansion     = ":::"
	linkedExpansion        = ":::+"
	cartesianFileExpansion = "::::"
	linkedFileExpansion    = "::::+"
)

// ArgumentColumnGroup - Stores column groups (appended using :::+)
// for the values being joined 1 to 1
type ArgumentColumnGroup struct {
	columns [][]string
}

// NewArgumentColumnGroup - Creates an empty ArgumentColumnGroup
func NewArgumentColumnGroup() *ArgumentColumnGroup {
	return &ArgumentColumnGroup{columns: [][]string{}}
}

// Append - Appends a value to a specific column while
// expanding the array as much as it's needed
func (g *ArgumentColumnGroup) Append(column int, value string) {
	for column >= len(g.columns) {
		g.columns = append(g.columns, []string{})
	}

	g.columns[column] = append(g.columns[column], value)
}

// Length - Return the maximum length available
// (which is the minimum of all lengths)
func (g *ArgumentColumnGroup) Length() int {
	min := len(g.columns[0])

	for _, vals := range g.columns[1:] {
		if len(vals) < min {
			min = len(vals)
		}
	}

	return min
}

// GetRow - Gets a row of 1 to 1 paired values
func (g *ArgumentColumnGroup) GetRow(rowNum int) []string {
	row := []string{}

	for _, column := range g.columns {
		row = append(row, column[rowNum])
	}

	return row
}

// ArgumentExpander - Expands arguments in a GNU parallel compatible way
type ArgumentExpander struct {
	current      int
	columnGroups []*ArgumentColumnGroup
}

// ParseArgumentExpansion - Parses a GNU parallel compatible argument list
// and returns an ArgumentExpander that behaves in a compatible way
func ParseArgumentExpansion(args []string) *ArgumentExpander {
	groups := []*ArgumentColumnGroup{}
	freader := fileReader{}

	currentGroup := 0
	currentColumn := 0
	fileExpansion := len(args) > 0 && (args[0] == cartesianFileExpansion || args[0] == linkedFileExpansion)
	expandGroup := len(args) > 0 && args[0] == cartesianFileExpansion

	for idx, arg := range args[1:] {
		if arg == cartesianExpansion {
			fileExpansion = false
			currentGroup++
		} else if arg == linkedExpansion {
			fileExpansion = false
			currentColumn++
		} else if arg == cartesianFileExpansion {
			fileExpansion = true
			expandGroup = true
		} else if arg == linkedFileExpansion {
			fileExpansion = true
			expandGroup = false
		} else {
			if fileExpansion && idx > 0 {
				if expandGroup {
					currentGroup++
				} else {
					currentColumn++
				}
			}

			if currentGroup >= len(groups) {
				groups = append(groups, NewArgumentColumnGroup())
				currentColumn = 0
			}

			if fileExpansion {
				for _, line := range freader.ReadLinesFromFile(arg) {
					groups[currentGroup].Append(currentColumn, line)
				}
			} else {
				groups[currentGroup].Append(currentColumn, arg)
			}
		}
	}

	return &ArgumentExpander{
		current:      -1,
		columnGroups: groups,
	}
}

// Length - Returns the number of possible rows in this expander
func (e *ArgumentExpander) Length() int {
	rowCount := 1

	for _, group := range e.columnGroups {
		rowCount *= group.Length()
	}

	return rowCount
}

// Next - Iterates to the next row and returns true if there's still more items
// Usage:
// for e.Next() {
//   e.Value()
// }
func (e *ArgumentExpander) Next() bool {
	e.current++
	return e.current < e.Length()
}

// Value - Returns the row consists of the current iteration
func (e *ArgumentExpander) Value() []string {
	values := []string{}

	div := 1

	for groupIdx := len(e.columnGroups) - 1; groupIdx >= 0; groupIdx-- {
		mod := e.columnGroups[groupIdx].Length()
		row := e.columnGroups[groupIdx].GetRow((e.current / div) % mod)

		values = append(row, values...)

		div *= e.columnGroups[groupIdx].Length()
	}

	return values
}
