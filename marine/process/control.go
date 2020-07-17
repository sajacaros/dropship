package process

import (
	"errors"
	"github.com/Masterminds/semver"
	"github.com/shirou/gopsutil/process"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type pidMap map[string]int
var pm = pidMap{}
var workingDir = "/var/local/vms/deploy"

var profile = "-Dspring.profiles.active=localization"

var versionRegex = regexp.MustCompile("[0-9]+")

func Start(project string) error {
	if err := checkByPidMap(project); err != nil {
		return err
	}

	if err := checkByProcessName(project); err != nil {
		return err
	}

	projectDir := projectDir(project)
	fullPath, err := jarFile(projectDir)
	if err != nil {
		return err
	}
	log.Println("start ----", fullPath, "----")
	cmd := exec.Command("java", "-Xmx512m", "-Xms256m", profile, "-jar", fullPath)
	cmd.Dir = projectDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	pm[project] = cmd.Process.Pid
	return nil
}

func Stop(project string) error {
	pid, exists := pm[project]

	if !exists {
		return errors.New("failed to retrieve pid, project : "+project)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Println(err)
		return errors.New("failed to find process, project : "+project+", pid : " + string(pid))
	}
	// Kill the process
	err = proc.Kill()
	if err != nil {
		log.Println(err)
		return errors.New("failed to kill process, project : "+project+", pid : " + string(pid))
	}

	delete(pm, project)

	return nil
}

func Update(project string) error {
	if err := Stop(project); err != nil {
		log.Println(err)
	}

	if err := Start(project); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func checkByPidMap(project string) error {
	if pid, existsInMap := pm[project]; existsInMap {
		existsInProcess, err := process.PidExists(int32(pid))
		if err != nil {
			return errors.New("failed to check that existsInProcess pid.., project : " + project + ", pid : " + string(pid))
		}
		if existsInProcess {
			return errors.New("already started.., please start after stopping or updating, project : " + project + ", pid : " + string(pid))
		} else {
			delete(pm, project)
		}
	}
	return nil
}

func checkByProcessName(project string) error {
	processes, _ := process.Processes()
	for _, pr := range processes {
		var cmdLine string
		var err error
		cmdLine, err = pr.Cmdline()
		if err == nil && strings.Contains(cmdLine, project) {
			return errors.New("already started.., please start after kill process, project : " + project + ", pid : " + string(pr.Pid))
		}
	}
	return nil
}

func jarFile(projectDir string) (string, error) {
	d, err := os.Open(projectDir)
	if err != nil {
		return "", errors.New("failed to open directory, dir : " + projectDir)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return "", errors.New("failed to read directory, dir : " + projectDir)
	}
	//^my-jar(\-\d+|\-\d+\.\d+)\.jar$
	targetFile := latestJarFile(files)
	if targetFile != "" {
		return projectDir + "/" + targetFile, nil
	} else {
		return "", errors.New("failed to find jar")
	}
}

func projectDir(project string) string {
	return workingDir + "/" + project
}

func latestJarFile(files []os.FileInfo) string {
	var targetFile string
	var prevVersion *semver.Version
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".jar" {
			continue
		}

		res := versionRegex.FindAllString(file.Name(), -1)
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
	return targetFile
}
