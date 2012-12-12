package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
}

var (
	cfg *Config = nil
)

func init() {
	bts, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("failed: Read config.json")
		return
	}

	cfg = new(Config)

	err = json.Unmarshal(bts, cfg)
	if err != nil {
		log.Fatal("failed: decode config.json")
		return
	}
}
