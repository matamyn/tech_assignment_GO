package main

import (
	"flag"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/app"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/link_shorter.yml", "path to config file")
}
func main() {
	flag.Parse()

	conf, err := link_shorter.ReadConf(configPath)
	if err != nil {
		log.Fatal(err)
	}
	app := link_shorter.New(conf)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
