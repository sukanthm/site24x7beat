package config

import "time"

type Config struct {
        Period time.Duration `config:"period"`
        Input []map[string]interface{} `config:"Input"`
}

var DefaultConfig = Config{
        Period: 3 * time.Second,
}
