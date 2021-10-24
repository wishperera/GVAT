package euvies

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	envKeyBaseURL    = "EU_VIES_BASE_URL"
	envKeyTimeout    = "EU_VIES_TIMEOUT"
	envKeyMaxRetries = "EU_VIES_MAX_RETRIES"

	defaultTimeoutSeconds = 5
	defaultMaxRetries     = 2

	failedToParseEnvKeyDue = "failed to parse env key: [%s] due: [%s]"
)

type Config struct {
	URL        string        `json:"url" env:"EU_VIES_BASE_URL"`
	Timeout    time.Duration `json:"timeout" env:"EU_VIES_TIMEOUT"`
	MaxRetries int           `json:"max_retries" env:"EU_VIES_MAX_RETRIES"`
}

func (c *Config) Init() error {
	c.URL = os.Getenv(envKeyBaseURL)
	ts := os.Getenv(envKeyTimeout)
	if ts == "" {
		c.Timeout = time.Second * defaultTimeoutSeconds
	} else {
		t, err := strconv.Atoi(ts)
		if err != nil {
			return fmt.Errorf(failedToParseEnvKeyDue, envKeyTimeout, err)
		}
		c.Timeout = time.Second * time.Duration(t)
	}

	mrs := os.Getenv(envKeyTimeout)
	if mrs == "" {
		c.MaxRetries = defaultMaxRetries
	} else {
		mr, err := strconv.Atoi(ts)
		if err != nil {
			return fmt.Errorf(failedToParseEnvKeyDue, envKeyTimeout, err)
		}
		c.MaxRetries = mr
	}

	return nil
}

func (c *Config) Print() string {
	byt, _ := json.Marshal(*c)
	return string(byt)
}

func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("%s not set", envKeyBaseURL)
	}

	if c.Timeout < 0 {
		return fmt.Errorf("%s invalid", envKeyTimeout)
	}

	if c.MaxRetries < 0 {
		return fmt.Errorf("%s invalid", envKeyMaxRetries)
	}

	return nil
}