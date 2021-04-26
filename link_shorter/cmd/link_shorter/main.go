package main

import (
	"flag"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"log"
	"runtime"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/link_shorter.yml", "path to config file")
}
func main() {
	flag.Parse()

	conf, err := common.ReadConf(configPath)
	if err != nil {
		log.Fatal(err)
	}
	runtime.GOMAXPROCS(2)
	if err := link_shorter.Start(conf); err != nil {
		log.Fatal(err)
	}
}
