package euvies

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	envKeyBaseURL          = "EU_VIES_BASE_URL"
	envKeyTimeout          = "EU_VIES_TIMEOUT"
	envKeyMaxRetries       = "EU_VIES_MAX_RETRIES"
	envKeyMaxWorkers       = "POOL_MAX_WORKERS"
	envKeyQueueSize        = "POOL_QUEUE_SIZE"
	envKeyWorkerBufferSize = "POOL_WORKER_BUFFER"

	defaultTimeoutSeconds   = 5
	defaultMaxRetries       = 2
	defaultMaxWorkers       = 10
	defaultQueueSize        = 1000
	defaultWorkerBufferSize = 10

	failedToParseEnvKeyDue = "failed to parse env key: [%s] due: [%s]"
)

type Config struct {
	URL          string        `json:"url" env:"EU_VIES_BASE_URL"`
	Timeout      time.Duration `json:"timeout" env:"EU_VIES_TIMEOUT"`
	MaxRetries   int           `json:"max_retries" env:"EU_VIES_MAX_RETRIES"`
	MaxWorkers   int           `json:"max_workers" env:"POOL_MAX_WORKERS"`
	QueueSize    int           `json:"queue_size" env:"POOL_QUEUE_SIZE"`
	WorkerBuffer int           `json:"worker_buffer" env:"POOL_WORKER_BUFFER"`
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

	maxWorkerStr := os.Getenv(envKeyMaxWorkers)
	if maxWorkerStr == "" {
		c.MaxWorkers = defaultMaxWorkers
	} else {
		maxWorkers, err := strconv.Atoi(maxWorkerStr)
		if err != nil {
			return fmt.Errorf(failedToParseEnvKeyDue, envKeyMaxWorkers, err)
		}
		c.MaxWorkers = maxWorkers
	}

	poolBufferSizeStr := os.Getenv(envKeyQueueSize)
	if poolBufferSizeStr == "" {
		c.QueueSize = defaultQueueSize
	} else {
		poolBufferSize, err := strconv.Atoi(poolBufferSizeStr)
		if err != nil {
			return fmt.Errorf(failedToParseEnvKeyDue, envKeyQueueSize, err)
		}
		c.QueueSize = poolBufferSize
	}

	workerBufferSizeStr := os.Getenv(envKeyWorkerBufferSize)
	if workerBufferSizeStr == "" {
		c.QueueSize = defaultWorkerBufferSize
	} else {
		pbs, err := strconv.Atoi(workerBufferSizeStr)
		if err != nil {
			return fmt.Errorf(failedToParseEnvKeyDue, envKeyWorkerBufferSize, err)
		}
		c.QueueSize = pbs
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
