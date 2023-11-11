package configuration

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

var CONFIG *Configuration

type Database struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Name       string `yaml:"name"`
	Engine     string `yaml:"engine"`
	SqlitePath string `yaml:"sqlitepath"`
	Debug      bool   `yaml:"debug"`
	DebugUser  bool   `yaml:"DebugUser"`
}

type Server struct {
	Motd           string  `yaml:"motd"`
	Version        float32 `yaml:"version"`
	Api            string  `yaml:"api"`
	Debug          bool    `taml:"debug"`
	Ip             string  `yaml:"ip"`
	Port           string  `yaml:"port"`
	LengthSecurity int     `yaml:"LengthSecurity"`
}

type Configuration struct {
	DB  Database `yaml:"Database"`
	Ser Server   `yaml:"Server"`
}

func Load(filename string) error {
	config := Configuration{}
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}
	CONFIG = &config
	//Cargo la configuracion
	if CONFIG.Ser.Debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.SetPrefix("DEBUG: ")
	}

	return nil
}
