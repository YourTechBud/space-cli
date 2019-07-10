package config

import (
	"io/ioutil"
	"encoding/json"

	"gopkg.in/yaml.v2"
)

// LoadConfigFromFile loads the config from the provided file path
func LoadConfigFromFile(path string) (*Deploy, error) {
	// Load the file in memory
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadConfigFromYAML(data)
}

// LoadConfigFromYAML loads the config from the provided yaml bytes
func LoadConfigFromYAML(text []byte) (*Deploy, error) {
	// Marshal the configuration
	conf := new(Deploy)
	err := yaml.Unmarshal(text, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// LoadConfigFromJSON loads the config from the provided yaml bytes
func LoadConfigFromJSON(text []byte) (*Deploy, error) {
	// Marshal the configuration
	conf := new(Deploy)
	err := json.Unmarshal(text, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
