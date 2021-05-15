package config

import (
	"time"
)

type Configuration struct {
	Debug  bool
	Client BaseClientConfiguration
	Social SocialConfiguration
}

type SocialConfiguration struct {
	ClientConfiguration `yaml:"httpClient,omitempty"`
	Concurrent bool
}

type BaseClientConfiguration struct{
	ClientConfiguration
}

type ClientConfiguration struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	MaxConnsPerHost     int
	IdleConnTimeout     time.Duration
}