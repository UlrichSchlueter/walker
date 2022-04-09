package exporter

import (
	"fmt"

	"github.com/ulrichschlueter/walker/pkg/utils/ntree"
)

type TreeExporter struct {
	Name    string
	Columns []string
}

func (t *TreeExporter) Export(tree *ntree.NTree, columns []string) {
	fmt.Println("['Location', 'Parent', 'Market trade volume (size)', 'Market increase/decrease (color)'],")
	t.preOrder(tree.Root, tree.Root.Level, 0)
}

//['Location', 'Parent', 'Market trade volume (size)', 'Market increase/decrease (color)'],
//['Global',    null,                 0,                               0],
//['America',   'Global',             0,                               0],
//['Europe',    'Global',             0,                               0],
//['Asia',      'Global',             0,                               0],
//['Australia', 'Global',             0,                               0],
//['Africa',    'Global',             0,                               0],
//['Brazil',    'America',            11,                              10],
func (t *TreeExporter) preOrder(node *ntree.TreeNode, path string, i int) {
	if node == nil {
		return
	}

	c := node.GetChildren()
	for _, k := range c {
		np := fmt.Sprintf("%s.%s", path, k.Level)

		fmt.Printf("['%s','%s',%d,%d],\n", np, path, len(k.Children)+1, 0)
		t.preOrder(k, np, i+1+len(c))

	}
}
