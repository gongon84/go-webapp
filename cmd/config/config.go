package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Env struct {
	YahooApiKey string `yaml:"yahoo_appid"`
}

type Configs struct {
	Env *Env
}

func ApplyConfig() Configs {
	configs := Configs{}

	e, err := applyEnv()
	if err != nil {
		log.Fatal(err)
	}

	configs.Env = &e

	return configs
}

func applyEnv() (Env, error) {
	buf, err := ioutil.ReadFile("./env.yaml")
	if err != nil {
		return Env{}, err
	}

	var e Env
	err = yaml.Unmarshal(buf, &e)

	return e, nil
}
