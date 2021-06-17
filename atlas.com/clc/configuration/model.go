package configuration

type Configuration struct {
	TimeoutTaskInterval int64 `yaml:"timeoutTaskInterval"`
	TimeoutDuration     int64 `yaml:"timeoutDuration"`
}
