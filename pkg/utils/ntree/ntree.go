package ntree

import (
	"fmt"
	"sort"
	"strings"
)

type Path struct {
	separator string
	split     []string
}

func NewPath(p string) *Path {
	return NewPathWithSep(p, ".")

}

func NewPathWithSep(p string, s string) *Path {

	path := &Path{
		separator: s,
		split:     strings.Split(p, s),
	}
	return path
}

func (p *Path) oneDown() *Path {
	return &Path{
		separator: p.separator,
		split:     p.split[1:],
	}
}

type TreeNode struct {
	Children map[string]*TreeNode
	sorted   []*TreeNode
	Level    string
	Payload  interface{}
}

type NTree struct {
	Root *TreeNode
}

func GetNAryTree() *NTree {
	// return new NAryTree
	return &NTree{Root: NewTreeNode("")}
}

func NewTreeNode(value string) *TreeNode {
	return &TreeNode{
		Payload:  nil,
		Level:    value,
		Children: make(map[string]*TreeNode),
		sorted:   nil,
	}
}

func (t *NTree) EnsurePath(path string) *TreeNode {
	p := *NewPath(path)

	return t.Root.ensurePath(p)
}

func (t *TreeNode) ensurePath(p Path) *TreeNode {
	if len(p.split) == 0 {
		return t
	}
	head := p.split[0]

	if _, ok := t.Children[head]; !ok {
		n := NewTreeNode(head)
		t.addChild(n)
	}

	child := t.Children[head]
	return child.ensurePath(*p.oneDown())

}

func (t *TreeNode) addChild(n *TreeNode) {
	t.Children[n.Level] = n
	t.sorted = make([]*TreeNode, 0)

	keys := make([]string, 0, len(t.Children))
	for k := range t.Children {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		t.sorted = append(t.sorted, t.Children[k])
	}
}

func (t *TreeNode) GetChildren() []*TreeNode {

	return t.sorted
}

func (t *NTree) PreOrder(node *TreeNode, level int) {
	if node == nil {
		return
	}
	var i int = 0
	var temp *TreeNode = nil
	if node.Payload == nil {
		fmt.Println(strings.Repeat("-", level), node.Level)
	} else {
		fmt.Println(strings.Repeat("-", level), node.Level, node.Payload)
	}
	// iterating the child of given node
	c := node.GetChildren()
	for i < len(c) {
		temp = c[i]
		t.PreOrder(temp, level+1)
		i++
	}
}
