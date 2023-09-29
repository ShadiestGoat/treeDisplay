package treeDisplay

import "strings"

// Renders the tree into a string.
func (c Config) RenderString(n Node) string {
	return strings.Join(c.RenderLines(n), "\n")
}

// Render the tree into a set of lines. This is useful mostly for internal use.
func (c Config) RenderLines(n Node) []string {
	children := n.GetChildren()

	lines := []string{}

	l1 := []rune{c.ESW, c.EW, ' '}

	if c.TrueTreeBranches {
		l1 = []rune{}
	} else if len(children) == 0 {
		l1[0] = c.EW
	}

	lines = append(lines, string(l1)+n.GetName())
	truePrefix := ""

	if c.TrueTreeBranches {
		truePrefix = " "

		if c.TrueTreeSpacing && len(n.GetName()) >= 2 {
			truePrefix += " "
		}
	}

	for i, ch := range children {
		startC := c.NES

		isLast := i == len(children)-1

		if isLast {
			startC = c.NE
		}

		cLines := c.RenderLines(ch)

		lines = append(lines, string([]rune{startC, c.EW})+cLines[0])

		prefix := string(c.NS) + " "

		if isLast {
			prefix = "  "
		}

		prefix += truePrefix

		for _, l := range cLines[1:] {
			lines = append(lines, prefix+l)
		}
	}

	return lines
}
