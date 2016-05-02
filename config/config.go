package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Redis  RedisConfig  `yaml:"redis"`
	Fluent FluentConfig `yaml:"fluent"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type FluentConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func Read() *Config {
	// Read config file
	buf, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	// Unmarshal yml
	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
