package process

import (
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"github.com/sajacaros/dropship/marine/config"
)

func Dependency() (*marine.ProjectDependency, error) {
	if dep, err := config.Dependency(); err!=nil {
		return nil, err
	} else {
		return &marine.ProjectDependency{Dependency: dep}, nil
	}
}
