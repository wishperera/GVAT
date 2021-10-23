package log

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const (
	defaultLogLevel = logLevelTrace

	envLogLevel        = "LOG_LEVEL"
	envFilePathEnabled = "LOG_FILE_PATH_ENABLED"
	envOutputDirectory = "OUTPUT_DIRECTORY"
)

type Config struct {
	Level           logLevel `json:"level" env:"LOG_LEVEL"`
	FilePathEnabled bool     `json:"file_path_enabled" env:"LOG_FILE_PATH_ENABLED"`
	OutputDirectory string   `json:"output_directory" env:"OUTPUT_DIRECTORY"`
}

func (c *Config) Init() (err error) {
	level := os.Getenv(envLogLevel)
	if level == "" {
		c.Level = defaultLogLevel
	} else {
		c.Level, err = parseLogLevel(level)
		if err != nil {
			return err
		}
	}

	fpe := os.Getenv(envFilePathEnabled)
	if fpe == "" {
		c.FilePathEnabled = true
	} else {
		c.FilePathEnabled, err = strconv.ParseBool(fpe)
		if err != nil {
			return fmt.Errorf("failed to parse env: %s due: %s", envFilePathEnabled, err)
		}
	}

	c.OutputDirectory = os.Getenv(envOutputDirectory)

	return nil
}

func (c *Config) Print() string {
	byt, _ := json.Marshal(*c)
	return string(byt)
}

func (c *Config) Validate() error {
	switch c.Level {
	case logLevelTrace, logLevelDebug, logLevelFatal, logLevelInfo, logLevelError:
	default:
		return fmt.Errorf("invalid log level: %s", c.Level)
	}

	return nil
}
