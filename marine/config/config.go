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
	Profile string
	Dependency string
	WorkingDirectory string `yaml:"workingDirectory"`
	AllowCommand string `yaml:"allowCommand"`
	CompletedMessage string `yaml:"completedMessage"`
	WaitTime int `yaml:"waitTime"`
}

var home string
var yamlPath string

func init() {
	home, _ = os.UserHomeDir()
	yamlPath = home + "/.dropship/config.yml"
}

func Source() (string, error) {
	options, err := yamlContents()
	if err != nil {
		return "", err
	}
	return options.Source, nil
}

func Dependency()(string, error) {
	options, err := yamlContents()
	if err != nil {
		return "", err
	}
	return options.Dependency, nil
}

func Profile() (string, error) {
	options, err := yamlContents()
	if err != nil {
		return "", err
	}
	return options.Profile, nil
}

func WaitTime() (int, error) {
	options, err := yamlContents()
	if err != nil {
		return -1, err
	}
	return options.WaitTime, nil
}

func WorkingDirectory() (string, error) {
	options, err := yamlContents()
	if err != nil {
		return "", err
	}
	return options.WorkingDirectory, nil
}

func Projects() ([]string, error) {
	options, err := yamlContents()
	if err != nil {
		return nil, err
	}
	return options.Projects, nil
}

func AllowCommand() (string, error) {
	options, err := yamlContents()
	if err != nil {
		return "", err
	}
	return options.AllowCommand, nil
}

func CompletedMessage() (string, error) {
	options, err := yamlContents()
	if err != nil {
		return "", err
	}
	return options.CompletedMessage, nil
}

func yamlContents() (*options, error) {
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
	return &options, nil
}
