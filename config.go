package main

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	RabbitMQURL string `toml:"rabbitmq_url"`
	SMTPHost    string `toml:"smtp_host"`
	SMTPPort    int    `toml:"smtp_port"`
	SMTPUser    string `toml:"smtp_username"`
	SMTPPass    string `toml:"smtp_password"`
	SMTPFrom    string `toml:"smtp_from"`
}

var cfg Config
var k = koanf.New(".")

func LoadConfig() {
	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
}

func GetConfig() Config {
	return cfg
}
