package process

import (
	"errors"
	"github.com/Masterminds/semver"
	"github.com/sajacaros/dropship/marine/config"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type pidMap map[string]int
var pm = pidMap{}
var workingDir = "/var/local/vms/deploy"

var profile = "-Dspring.profiles.active=localization"

var versionRegex = regexp.MustCompile("[0-9]+")

func Start(project string) error {
	if err := checkRunningByPidMap(project); err != nil {
		return err
	}

	if err := checkRunningByName(project); err != nil {
		return err
	}

	projectDir := projectDir(project)
	fileName, err := findJarFileName(projectDir, project)
	fullPath := projectDir+"/"+fileName
	if err != nil {
		return err
	}
	log.Println("start ----", fullPath, "----")
	cmd := exec.Command("java", "-Xmx512m", "-Xms256m", profile, "-jar", fullPath)
	cmd.Dir = projectDir
	err = cmd.Start()
	pm[project] = cmd.Process.Pid
	return nil
}

func Stop(project string) error {
	pid, err := findPidByMap(project)
	if err != nil {
		pid, err = findPidByName(project)
		if err != nil { // process가 존재하지 않음
			log.Println(project, "가 존재하지 않음")
			return nil
		}
	}

	log.Println("pid : ", strconv.Itoa(pid))

	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		log.Println(err)
		return errors.New("failed to find process, project : "+project+", pid : " + strconv.Itoa(pid))
	}

	// Kill the process
	err = proc.Kill()
	if err != nil {
		log.Println(err)
		return errors.New("failed to kill process, project : "+project+", pid : " + strconv.Itoa(pid))
	}

	delete(pm, project)

	return nil
}

func Update(project string) error {
	source, err := config.Source()
	if err != nil {
		return err
	}
	err = jarFileCopy(source, project)
	if err != nil {
		return err
	}

	if err = Stop(project); err != nil {
		log.Println(err)
	}

	if err = Start(project); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func Install() error {
	source, err := config.Source()
	if err != nil {
		return err
	}
	projects, err := config.Projects()
	if err != nil {
		return err
	} else if len(projects) < 1 {
		return errors.New("empty project")
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(projects))
	for _, project := range projects {
		go func(p string) {
			err = jarFileCopy(source, p)
			if err!=nil {
				log.Println("copy error : ", err)
			}
			waitGroup.Done()
		}(project)
	}

	go func() {
		waitGroup.Wait()
	}()

	return nil
}



func jarFileCopy(source string, project string) error {
	fileName, err := findJarFileName(source, project)
	if err != nil {
		return err
	}

	err = copy(source+"/"+fileName, projectDir(project)+"/"+fileName)
	if err != nil {
		return err
	}
	return nil
}

func copy(src string, dst string) error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	if err!=nil {
		return err
	}
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	if err!=nil {
		return err
	}
	return nil
}

func findPidByMap(project string) (int, error){
	pid, exists := pm[project]
	if !exists {
		return -1, errors.New("failed to retrieve pid from map, project : "+project)
	}

	if exists, err := process.PidExists(int32(pid)); err != nil || !exists {
		delete(pm, project)
		return -1, errors.New("not exist running process, project : "+project+", pid : " + strconv.Itoa(pid))
	}
	return pid, nil
}

func checkRunningByPidMap(project string) error {
	if pid, existsInMap := pm[project]; existsInMap {
		existsInProcess, err := process.PidExists(int32(pid))
		if err != nil {
			return errors.New("failed to check that existsInProcess pid.., project : " + project + ", pid : " + strconv.Itoa(pid))
		}
		if existsInProcess {
			return errors.New("already started.., please start after stopping or updating, project : " + project + ", pid : " + strconv.Itoa(pid))
		} else {
			delete(pm, project)
		}
	}
	return nil
}

func checkRunningByName(project string) error {
	processes, _ := process.Processes()
	for _, pr := range processes {
		var cmdLine string
		var err error
		cmdLine, err = pr.Cmdline()
		if err == nil && strings.Contains(cmdLine, project) {
			return errors.New("already started.., please start after kill process, project : " + project + ", pid : " + strconv.Itoa(int(pr.Pid)))
		}
	}
	return nil
}

func findPidByName(project string) (int, error) {
	processes, _ := process.Processes()
	for _, pr := range processes {
		var cmdLine string
		var err error
		cmdLine, err = pr.Cmdline()
		if err == nil && strings.Contains(cmdLine, project) {
			return int(pr.Pid), nil
		}
	}
	return -1, errors.New("not exist running project, project : " + project)
}

func findJarFileName(projectDir string, project string) (string, error) {
	files, err := ioutil.ReadDir(projectDir)
	if err != nil {
		return "", errors.New("failed to read directory, dir : " + projectDir)
	}
	//^my-jar(\-\d+|\-\d+\.\d+)\.jar$
	targetFile := latestJarFile(files, project)
	if targetFile != "" {
		return targetFile, nil
	} else {
		return "", errors.New("failed to find jar")
	}
}

func projectDir(project string) string {
	return workingDir + "/" + project
}

func latestJarFile(files []os.FileInfo, project string) string {
	var targetFile string
	var prevVersion *semver.Version
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".jar" {
			continue
		}

		if match, err := regexp.MatchString(project, file.Name()); !match || err!=nil {
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
