package model

import (
	"strings"
)

type Config struct {
	data map[string]interface{}
}

func (conf *Config) GetValueForKey(key string) (interface{}, bool) {
	val, ok := conf.data[strings.ToLower(key)]
	return val, ok
}

func NewConfig(data map[string]interface{}) *Config {
	return &Config{data: data}
}
