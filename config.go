package treeDisplay

type Config struct {
	// East, South, West
	// Indicates that the node has children, and connects to them
	ESW,
	// North, East, South
	// Used as the split between a vertical line
	NES,

	// East, West (Horizontal)
	EW,
	// North, South (Vertical)
	NS,

	// North, East
	// Indicates that this is the last child, and doesn't continue the vertical line.
	NE rune

	// If this is true, it won't extend the tree branches, and render like this:
	// Parent
	// ├─ Child
	// ╰─ Other CHild
	TrueTreeBranches bool
	// If enabled, with TrueTreeBranches enabled, this will add an extra space if possible to the child branches
	TrueTreeSpacing bool
}

var ConfigRounded = Config{
	ESW: '┬',
	NES: '├',

	EW: '─',
	NS: '│',

	NE: '╰',
}

var ConfigStraight = Config{
	ESW: '┬',
	NES: '├',

	EW: '─',
	NS: '│',

	NE: '└',
}

var ConfigBold = Config{
	ESW: '┳',
	NES: '┣',

	EW: '━',
	NS: '┃',

	NE: '┗',
}

var ConfigDouble = Config{
	ESW: '╦',
	NES: '╠',

	EW: '═',
	NS: '║',

	NE: '╚',
}
