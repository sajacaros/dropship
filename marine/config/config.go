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
	WorkingDir string
	AllowCommand string
	CompletedMessage string
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

func WorkingDir() (string, error) {
	options, err := yamlContents()
	if err != nil {
		return "", err
	}
	return options.WorkingDir, nil
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
