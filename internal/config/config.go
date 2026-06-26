package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

const (
	defaultHttpPort           = 50000
	defaultHttpReadTimeout    = 5
	defaultHttpWriteTimeout   = 10
	defaultHttpIdleTimeout    = 60
	defaultHttpShutdownTimout = 30
	defaultSnowflakeNodeId    = 1
)

type HttpConfig struct {
	Host            string `koanf:"host"`
	Port            int    `koanf:"port"`
	ReadTimeout     int    `koanf:"read_timeout"`
	WriteTimeout    int    `koanf:"write_timeout"`
	IdleTimeout     int    `koanf:"idle_timeout"`
	ShutdownTimeout int    `koanf:"shutdown_timeout"`
}

type PostgresConfig struct {
	ConnString string `koanf:"conn_string"`
}

type SnowflakeConfig struct {
	NodeID int64 `koanf:"node_id"`
}

type Config struct {
	Http      HttpConfig      `koanf:"http"`
	Postgres  PostgresConfig  `koanf:"postgres"`
	Snowflake SnowflakeConfig `koanf:"snowflake"`
}

func Load() (*Config, error) {
	k := koanf.New(".")

	if err := k.Load(confmap.Provider(map[string]any{
		"http.host":             "0.0.0.0",
		"http.port":             defaultHttpPort,
		"http.read_timeout":     defaultHttpReadTimeout,
		"http.write_timeout":    defaultHttpWriteTimeout,
		"http.idle_timeout":     defaultHttpIdleTimeout,
		"http.shutdown_timeout": defaultHttpShutdownTimout,
		"snowflake.node_id":     defaultSnowflakeNodeId,
	}, "."), nil); err != nil {
		return nil, fmt.Errorf("load defaults: %w", err)
	}

	path := "config/config.yaml"
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("load config file: %w", err)
	}

	if err := k.Load(env.Provider("GM_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "GM_")
		s = strings.ToLower(s)
		s = strings.ReplaceAll(s, "__", ".")

		return s
	}), nil); err != nil {
		return nil, fmt.Errorf("load env: %w", err)
	}

	cfg := &Config{}
	if err := k.Unmarshal("", cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return cfg, nil
}
