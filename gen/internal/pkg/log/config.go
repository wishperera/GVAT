package log

type Config struct {
	Level           string `env:"LOG_LEVEL"`
	FilePathEnabled bool   `env:"LOG_FILE_PATH_ENABLED"`
	OutputDirectory string `env:"OUTPUT_DIRECTORY"`
}

func (c *Config) Init() error {

}

func (c *Config) Print() string {
	panic("implement me")
}

func (c *Config) Validate() error {
	panic("implement me")
}
