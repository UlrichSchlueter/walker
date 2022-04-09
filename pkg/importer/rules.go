package importer

import (
	"fmt"
	"io/ioutil"

	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type IgnoreRule struct {
	Name          string `json:"name"`
	IfPathMatches string `json:"ifpathmatches"`
}

type ReplaceRule struct {
	Name             string `json:"name"`
	IfPathMatches    string `json:"ifpathmatches"`
	ReplaceValueWith string `json:"replacevaluewith"`
	ReplacePathWith  string `json:"replacepathwith"`
}

type Rules struct {
	IgnoreRules  []IgnoreRule  `json:"ignore"`
	ReplaceRules []ReplaceRule `json:"replace"`
}

func (r *Rules) apply(path string, value string) (newpath string, newvalue string, ignore bool) {

	newpath = path
	newvalue = value
	ignore = false

	for _, r := range r.IgnoreRules {
		if r.IfPathMatches != "" {
			matched, _ := regexp.Match(r.IfPathMatches, []byte(path))
			if matched {
				log.Debugf("Ignore rule '%s' matched path: %s", r.Name, path)
				ignore = true
				return
			}
		}
	}

	for _, r := range r.ReplaceRules {
		if r.IfPathMatches != "" {
			matched, _ := regexp.Match(r.IfPathMatches, []byte(path))
			if matched {
				log.Debugf("Replace rule '%s' matched path: %s", r.Name, path)

				newvalue = r.ReplaceValueWith
				return
			}
		}
	}

	return

}

func (r *Rules) writeRules(folder string) {
	s, _ := yaml.Marshal(r)

	err := os.WriteFile(folder+"/rules.yaml", []byte(s), 0644)
	if err != nil {
		fmt.Printf("Error writing Jobs. %v", err)
	}

}

func (r *Rules) readRules(filename string) {
	yfile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var f *Rules
	err2 := yaml.Unmarshal(yfile, &f)

	if err2 != nil {
		log.Fatal(err2)
	}
	r.IgnoreRules = f.IgnoreRules
	r.ReplaceRules = f.ReplaceRules

}
