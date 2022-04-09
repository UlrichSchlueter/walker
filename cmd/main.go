package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/ulrichschlueter/walker/pkg/collect"
	"github.com/ulrichschlueter/walker/pkg/compare"
)

func main() {
	var action string
	var folder string
	var exporter string
	var comparerules string
	var loglevel string
	flag.StringVar(&action, "a", "collect", "action")
	flag.StringVar(&folder, "f", "folder", "folder")
	flag.StringVar(&exporter, "e", "export", "exporter")
	flag.StringVar(&comparerules, "i", "", "rules")
	flag.StringVar(&loglevel, "loglevel", "Info", "log level")
	flag.Parse()

	lev, err := log.ParseLevel(loglevel)
	if err != nil {
		log.Fatalf("Invalid log level%s", loglevel)
	}
	log.SetLevel(lev)

	switch action {
	case "collect":
		collect.Collect(folder)
	case "compare":
		compare.Compare(folder, exporter, comparerules)
	default:
		log.Fatalf("unexpected action:%s", action)

	}

}
