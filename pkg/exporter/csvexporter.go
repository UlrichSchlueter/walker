package exporter

import (
	"fmt"

	"github.com/ulrichschlueter/walker/pkg/utils/ntree"
)

type CsvExporter struct {
	Name    string
	Columns []string
}

func (csv *CsvExporter) Export(tree *ntree.NTree, columns []string) {

	csv.Columns = columns

	cols := ""
	for _, c := range columns {
		cols = cols + "," + c
	}
	fmt.Println("path" + cols)
	csv.preOrder(tree.Root, "")
}

func (csv *CsvExporter) preOrder(node *ntree.TreeNode, path string) {
	if node == nil {
		return
	}

	if node.Payload != nil {
		colString := ","
		isFirstValue := true
		hasdifferentvalues := false
		ref := ""
		//l := ""
		for _, c := range csv.Columns {
			p := node.Payload.(map[string]string)

			if isFirstValue {
				ref = p[c]
				colString = p[c]
				isFirstValue = false
			} else {
				if ref == p[c] {
					colString = colString + ",<=="
				} else {
					colString = colString + "," + p[c]
					hasdifferentvalues = true
				}
			}
		}
		if hasdifferentvalues {
			fmt.Println(path + "." + node.Level + "," + colString)
		}

	}

	c := node.GetChildren()
	for _, k := range c {
		if node.Level == "" {
			csv.preOrder(k, path)
		} else {
			csv.preOrder(k, path+"."+node.Level)
		}
	}
}
