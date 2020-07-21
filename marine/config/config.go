package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type options struct {
	Source string
	Projects []string `yaml:",flow"`
}

var yamlFile = "$HOME/.dropship/config.yml"

func Source() (string, error) {
	filename, _ := filepath.Abs(yamlFile)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("failed to read yaml file")
	}

	options := options{}

	err = yaml.Unmarshal(yamlFile, &options)
	if err != nil {
		return "", errors.New("failed to unmarshal yaml file")
	}
	return options.Source, nil
}

func Projects() ([]string, error) {
	filename, _ := filepath.Abs(yamlFile)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("failed to read yaml file")
	}

	options := options{}

	err = yaml.Unmarshal(yamlFile, &options)
	if err != nil {
		return nil, errors.New("failed to unmarshal yaml file")
	}
	return options.Projects, nil
}