package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type ConfigModel struct {
	Database struct {
		User     string `yaml:"user,omitempty"`
		Password string `yaml:"password,omitempty"`
		Host     string `yaml:"host,omitempty"`
		Port     string `yaml:"port,omitempty"`
		Database string `yaml:"database,omitempty"`
	} `yaml:"database,omitempty"`

	Port struct {
		Port string `yaml:"port,omitempty"`
	} `yaml:"port,omitempty"`
}

func GetConfigs() ConfigModel {
	c, err := readConf("./common/config/local.yaml")
	if err != nil {
		log.Fatal(err)
	}

	return *c
}

func readConf(filename string) (*ConfigModel, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &ConfigModel{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}
