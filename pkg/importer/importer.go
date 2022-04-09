package importer

import (
	"encoding/json"
	"fmt"

	"sort"

	log "github.com/sirupsen/logrus"

	ntree "github.com/ulrichschlueter/walker/pkg/utils/ntree"
)

type lineData struct {
	Data map[string]string
}

type Importer struct {
	Columns map[string]bool
	Lines   map[string]lineData
	Tree    *ntree.NTree
	rules   *Rules
}

func NewImporter() *Importer {
	return &Importer{
		Columns: make(map[string]bool),
		Lines:   make(map[string]lineData),
		Tree:    ntree.GetNAryTree(),
		rules: &Rules{
			IgnoreRules:  make([]IgnoreRule, 0),
			ReplaceRules: make([]ReplaceRule, 0)},
	}

}

func (c Importer) getSortedListOfKeys(f map[string]interface{}) (keys []string) {
	// Get a list of sorted keys of all the children
	keys = make([]string, 0, len(f))
	for k := range f {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}

func (c Importer) processLevel(f map[string]interface{}, path string, name string) {
	if f == nil {
		return
	}

	// iterate over all the children by sorted key
	for _, k := range c.getSortedListOfKeys(f) {
		p := path + "." + k
		switch v := f[k].(type) {
		case []interface{}:
			for i, k := range v {
				pp := fmt.Sprintf("%s.%d", p, i)
				c.handleChild(k, pp, name)
			}
		default:
			c.handleChild(v, p, name)
		}

	}
}

func (c Importer) createOrUpdateLeaf(path string, name string, value string) {
	//insert into Tree
	node := c.Tree.EnsurePath(path)

	//c.Tree.PreOrder(c.Tree.Root, 0)
	var y map[string]string
	if node.Payload == nil {
		y = make(map[string]string)
	} else {
		y = node.Payload.(map[string]string)
	}
	y[name] = value
	c.Columns[name] = true
	node.Payload = y

	if _, ok := c.Lines[path]; !ok {
		c.Lines[path] = lineData{Data: make(map[string]string)}
	}
	c.Lines[path].Data[name] = value

}

func (c Importer) handleChild(x interface{}, path string, name string) {
	value := "<unknown>"
	switch v := x.(type) {
	case string, float64, bool, nil:
		value = fmt.Sprintf("%v", v)
	case map[string]interface{}:
		n := x.(map[string]interface{})
		c.processLevel(n, path, name)
	default:
		// if the source is json formatted data, this should be impossible.
		log.Panicf("Can't import data: unrecognized data type %v", v)
	}

	// must be leaf, apply ignore and replace regex rules, if any
	newp, newv, ignore := c.rules.apply(path, value)

	if !ignore {
		c.createOrUpdateLeaf(newp, name, newv)
	}

}

func (c Importer) ProcessDataSet(b []byte, meta []byte, folder string, rulesFile string) {

	var f map[string]interface{}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(meta, &m); err != nil {
		panic(err)
	}
	p := m["type"].(string)
	i := m["index"].(string)
	pre := m["prefix"].(string)

	if rulesFile != "" {
		c.rules.readRules(rulesFile)
	}

	c.processLevel(f, fmt.Sprintf("%s.%s.%s", pre, p, i), folder)

}
