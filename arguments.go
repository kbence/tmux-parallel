package main

type ArgumentExpander struct {
	current int
	values  []string
}

func ParseArgumentExpansion(args []string) *ArgumentExpander {
	return &ArgumentExpander{
		current: -1,
		values:  args,
	}
}

func (e *ArgumentExpander) Next() bool {
	e.current++
	return e.current < len(e.values)
}

func (e *ArgumentExpander) Value() string {
	return e.values[e.current]
}
