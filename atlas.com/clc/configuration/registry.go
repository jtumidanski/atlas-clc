package configuration

import (
	"sync"
)

type ConfigurationRegistry struct {
	c *Configuration
	e error
}

var configurationRegistryOnce sync.Once
var configurationRegistry *ConfigurationRegistry

func GetConfiguration() (*Configuration, error) {
	configurationRegistryOnce.Do(func() {
		configurationRegistry = &ConfigurationRegistry{}
		err := configurationRegistry.loadConfiguration()
		configurationRegistry.e = err
	})
	return configurationRegistry.c, configurationRegistry.e
}
