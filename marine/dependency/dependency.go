package dependency

import (
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/sajacaros/dropship/marine/config"
	"regexp"
	"strings"
)

type Unity interface {
	parent() (*Unity, error)
	hasProject(string) bool
}

type Node struct {
	Projects *[]string
	Parent *Node
}

var root Node
var current *Node

var r *regexp.Regexp

func init() {
	current = &root
	r = regexp.MustCompile(`[A-Za-z]+`)
}
func (node *Node)parent() (*Node, error) {
	if node.Parent == &root {
		return nil, errors.New("not exist parent")
	}
	return node.Parent, nil
}

func (node *Node) hasProject(project string) bool {
	for _, p := range *node.Projects {
		if strings.EqualFold(project,p) {
			return true
		}
	}
	return false
}

func findNode(node *Node, project string) (*Node, error) {
	if node == &root {
		return nil, errors.New("not exist")
	}

	if node.hasProject(project) {
		return node, nil
	}
	return findNode(node.Parent, project)
}

func ReadDependency(project string) (*[]string, error){
	current = &root
	dependencyStr,err := config.Dependency()
	if err != nil {
		return nil, err
	}

	makeLinkedList(dependencyStr)

	node, err := findNode(current, project)
	if err != nil {
		return drawDependencyInclude(current), nil
	}

	return drawDependencyExclude(node), nil
}

func drawDependencyInclude(node *Node) *[]string{
	return drawDependency(node, true)
}

func drawDependencyExclude(node *Node) *[]string{
	return drawDependency(node, false)
}

func drawDependency(node *Node, include bool) *[]string {
	var dependencyArr []string
	if include {
		dependencyArr = append(*node.Projects)
	}
	for node.Parent != &root {
		dependencyArr = append(*node.Parent.Projects, dependencyArr...)
		node = node.Parent
	}
	current = &root
	return &dependencyArr
}

func makeLinkedList(dependencyStr string) {
	dependencyArr := strings.Split(dependencyStr, "-")
	log.Info(dependencyArr)
	for _, dependency := range dependencyArr {
		var t Node
		if strings.HasPrefix(dependency, "[") {
			log.Info(dependency + " dependency")
			regexp.MustCompile(`\w`)
			project := r.FindAllString(dependency, -1)
			t = Node{Parent: current, Projects: &project}

		} else {
			log.Info(dependency + " Edge")
			t = Node{Parent: current, Projects: &[]string{dependency}}
		}
		current = &t
	}
}
