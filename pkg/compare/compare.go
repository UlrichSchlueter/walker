package compare

import (
	"os"
	"strings"

	"github.com/ulrichschlueter/walker/pkg/exporter"
	"github.com/ulrichschlueter/walker/pkg/importer"
)

func Compare(folders string, exporterType string, rulesFile string) {
	f := strings.Split(folders, ":")

	// import all folders and filter data
	c := importer.NewImporter()

	for _, folder := range f {
		f := getAllFilesFromFolder(folder)
		for _, file := range f {
			s := getJsonFromFile(folder + string(os.PathSeparator) + file)
			m := getJsonFromFile(folder + string(os.PathSeparator) + file + ".meta")
			c.ProcessDataSet(s, m, folder, rulesFile)
		}
	}

	cols := make([]string, 0, len(f))
	for k := range c.Columns {
		cols = append(cols, k)
	}

	// export all data
	var e exporter.Exporter
	switch exporterType {
	case "csv":
		e = &exporter.CsvExporter{}
	default:
		e = &exporter.TreeExporter{}
	}
	e.Export(c.Tree, cols)

}
