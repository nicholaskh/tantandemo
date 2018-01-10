package main

import (
	"flag"
)

var (
	options struct {
		configFile   string
	}
)

func parseFlags() {
	flag.StringVar(&options.configFile, "conf", "etc/test.cf", "config file")

	flag.Parse()
}
