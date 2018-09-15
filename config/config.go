package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"flag"
)

type Config struct {
	Database database
}


type database struct {
	Server string
	Port string
	Database string
	User string
	Password string
}

func (c *Config) Read() {

	var filename = flag.String("config", "config.toml", "Location of the config file.")
	flag.Parse()
	if _, err := toml.DecodeFile(*filename, &c); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
