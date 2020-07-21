package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type options struct {
	Source string
	Projects []string `yaml:",flow"`
}

var home string
var yamlPath string

func init() {
	home, _ = os.UserHomeDir()
	yamlPath = home + "/.dropship/config.yml"
}

func Source() (string, error) {
	filename, _ := filepath.Abs(yamlPath)
	yamlContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("failed to read yaml file")
	}

	options := options{}

	err = yaml.Unmarshal(yamlContents, &options)
	if err != nil {
		return "", errors.New("failed to unmarshal yaml file")
	}
	return options.Source, nil
}

func Projects() ([]string, error) {
	filename, _ := filepath.Abs(yamlPath)
	yamlContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("failed to read yaml file")
	}

	options := options{}

	err = yaml.Unmarshal(yamlContents, &options)
	if err != nil {
		return nil, errors.New("failed to unmarshal yaml file")
	}
	return options.Projects, nil
}