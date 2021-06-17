package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func (c *ConfigurationRegistry) loadConfiguration() error {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	con := &Configuration{}
	err = yaml.Unmarshal(yamlFile, con)
	if err != nil {
		return err
	}
	c.c = con
	return nil
}