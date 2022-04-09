package importer

import (
	"fmt"
	"testing"
)

func TestWriteRules(t *testing.T) {
	r := Rules{
		IgnoreRules: []IgnoreRule{
			{Name: "test",
				IfPathMatches: "n",
			},
		},
		ReplaceRules: []ReplaceRule{
			{Name: "re",
				IfPathMatches:   "a",
				ReplacePathWith: "b",
			},
		},
	}
	r.writeRules("../..")
}

func TestReadFile(t *testing.T) {
	var r = Rules{}
	r.readRules("../../testrules.yaml")
	fmt.Println(r)
}
