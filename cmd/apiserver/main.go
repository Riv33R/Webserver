package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/Riv33R/Webserver/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
