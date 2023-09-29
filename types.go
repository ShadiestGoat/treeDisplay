package treeDisplay

type Node interface {
	GetChildren() []Node
	GetName() string
}

type TreeNode struct {
	Name     string
	Children []Node
}

func (t TreeNode) GetChildren() []Node {
	return t.Children
}

func (t TreeNode) GetName() string {
	return t.Name
}
