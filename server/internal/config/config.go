package config

import (
	"bytes"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Listen string `yaml:"listen"`
}

func Default() Config {
	return Config{Listen: "127.0.0.1:8080"}
}

func Load(path string) (Config, error) {
	cfg := Default()
	if path == "" {
		return cfg, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse config %s: %w", path, err)
	}
	return cfg, nil
}
