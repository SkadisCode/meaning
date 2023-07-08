package config

import (
	"fmt"
	"io/ioutil"

	"github.com/SkadisCode/meaning/entity"
	"gopkg.in/yaml.v2"
)

func ReadConfig(filename string) (*entity.SQLConfig, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %v", err)
	}

	// Unmarshal the YAML data into the SQLConfig struct
	var sqlConfig entity.SQLConfig
	err = yaml.Unmarshal(data, &sqlConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML data: %v", err)
	}

	return &sqlConfig, nil
}
