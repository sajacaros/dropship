package process

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"github.com/Masterminds/semver"
	"strings"
)

type pidMap map[string]int
var pm = pidMap{}
var workingDir = "/var/local/vms/deploy/"

var command = "java"
var profile = "-Dspring.profiles.active=localization"

func Start(project string) error {
	dir := "/var/local/vms/deploy/"+project+"/"
	projectWithExt := project+".jar"
	fullPathWithExt := dir+projectWithExt

	d, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to open directory, dir : "+ dir)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to read directory, dir : "+ dir)
	}
	//^my-jar(\-\d+|\-\d+\.\d+)\.jar$
	m := regexp.MustCompile("[0-9]+")
	var targetFile string
	var prevVersion *semver.Version
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".jar" {
			continue
		}

		res := m.FindAllString(file.Name(), -1)
		version, err := semver.NewVersion(strings.Join(res, "."))
		if err != nil {
			log.Println("Failed to acquire version")
			continue
		}

		if prevVersion == nil {
			targetFile = file.Name()
			prevVersion = version
		} else {
			if prevVersion.Compare(version) < 0 {
				targetFile = file.Name()
				prevVersion = version
			}
		}

	}
	fmt.Println(targetFile)

	//m := regexp.MustCompile(`^my-jar(\-\d+|\-\d+\.\d+)\*.jar$`)
	//res := m.FindAllString("abcd-1.2.3-snp.jar", -1)
	fmt.Println(res)

	sort.Sort(sort.Reverse(sort.StringSlice(jarFiles)))
	fmt.Println(jarFiles)

	cmd := exec.Command(command, "-Xmx512m", "-Xms256m", profile, "-jar", fullPathWithExt)
	//cmd := exec.Command("ls", "-al")
	cmd.Dir = "/var/local/vms/deploy/extractor"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	return nil
}
func Stop(project string) error {
	log.Println("stop ", project, "~")
	log.Println("pid : ", pm["abc"])
	return nil
}
func Update(project string) error {
	return nil
}
