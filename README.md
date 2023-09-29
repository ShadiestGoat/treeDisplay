# Tree Display

[![Go Reference](https://pkg.go.dev/badge/github.com/shadiestgoat/treeDisplay.svg)](https://pkg.go.dev/github.com/shadiestgoat/treeDisplay) 

A simple and extendable go library to render tree-like elements into a string/string array

## Installing

`go get github.com/shadiestgoat/treeDisplay`

## How to use

Since we don't know the width of the tree, this library uses a recursive solution. As such, you need `Nodes`. A `Node` is an element of the tree, that has to follow this interface:

```go
type Node interface {
	GetChildren() []Node
	GetName() string
}
```

An example node is provided as `TreeNode`. You would then need to provide a `Config`. A config just provides runes which are displayed. Provided configs are named `Config*`. If you want to create your own, you can follow the example of those & comments on the `Config` type.

Then, to render your tree, use the `Config.RenderString(rootNode)` or `Config.RenderLines(rootNode)`.