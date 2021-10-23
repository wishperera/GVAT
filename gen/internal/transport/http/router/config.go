package router

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	envKeyServerPort      = "SERVER_PORT"
	envKeyReadTimeout     = "SERVER_READ_TIMEOUT"
	envKeyWriteTimeout    = "SERVER_WRITE_TIMEOUT"
	defaultPort           = "8889"
	defaultTimeoutSeconds = 10 // timeout in seconds

	invalidValueProvidedForEnv = "invalid value provided for env: %s"

	base10    = 10
	bitSize64 = 64
)

type Config struct {
	Port         string `env:"SERVER_PORT"`
	ReadTimeout  int64  `env:"SERVER_READ_TIMEOUT"`
	WriteTimeout int64  `env:"SERVER_WRITE_TIMEOUT"`
}

func (c *Config) Init() (err error) {
	c.Port = os.Getenv(envKeyServerPort)
	if c.Port == "" {
		c.Port = defaultPort
	}

	rt := os.Getenv(envKeyReadTimeout)
	if rt == "" {
		c.ReadTimeout = defaultTimeoutSeconds
	} else {
		c.ReadTimeout, err = strconv.ParseInt(rt, base10, bitSize64)
		if err != nil {
			return fmt.Errorf("failed to parse env: %s due: %s", envKeyReadTimeout, err)
		}
	}

	wt := os.Getenv(envKeyReadTimeout)
	if wt == "" {
		c.WriteTimeout = defaultTimeoutSeconds
	} else {
		c.WriteTimeout, err = strconv.ParseInt(rt, base10, bitSize64)
		if err != nil {
			return fmt.Errorf("failed to parse env: %s due: %s", envKeyWriteTimeout, err)
		}
	}

	return nil
}

func (c *Config) Print() string {
	return fmt.Sprintf("server_port: %s", c.Port)
}

func (c *Config) Validate() error {
	matched, err := regexp.Match("[[:digit:]]", []byte(c.Port))
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf(invalidValueProvidedForEnv, envKeyServerPort)
	}

	if c.ReadTimeout < 0 {
		return fmt.Errorf(invalidValueProvidedForEnv, envKeyReadTimeout)
	}

	if c.WriteTimeout < 0 {
		return fmt.Errorf(invalidValueProvidedForEnv, envKeyWriteTimeout)
	}

	return err
}
