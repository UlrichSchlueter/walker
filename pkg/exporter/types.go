package exporter

import "github.com/ulrichschlueter/walker/pkg/utils/ntree"

type Exporter interface {
	Export(tree *ntree.NTree, columns []string)
}
