package process

import (
	"bufio"
	"context"
	"errors"
	"github.com/Masterminds/semver"
	"github.com/hako/durafmt"
	"github.com/sajacaros/dropship/build/gen/bnpinnovation.com/marine"
	"github.com/sajacaros/dropship/marine/config"
	"github.com/sajacaros/dropship/marine/dependency"
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
	"time"
)

type pidMap map[string]int
var pm = pidMap{}

var profileBase = "-Dspring.profiles.active="

var versionRegex = regexp.MustCompile("[0-9]+")

func Start(project string) error {
	if err := assertDependency(project); err!=nil {
		return err
	}

	if err := checkRunningByPidMap(project); err != nil {
		return err
	}

	if err := checkRunningByName(project); err != nil {
		return err
	}

	projectDir := projectDir(project)
	fileName, err := findJarFileName(projectDir, project)
	if err != nil {
		log.Println("projectDir : ", projectDir)
		return err
	}
	fullPath := projectDir+"/"+fileName
	log.Println("start ----", fullPath, "----")

	var profile string
	profile, err = config.Profile()
	if err!=nil  {
		return err
	}

	cmd := exec.Command("java", "-Xmx512m", "-Xms256m", profileBase+profile, "-jar", fullPath)
	cmd.Dir = projectDir
	r, _ := cmd.StdoutPipe()

	completeChannel := make(chan struct{})
	watcherChannel := make(chan bool, 1)
	
	scanner := bufio.NewScanner(r)
	ctx, cancel := context.WithCancel(context.Background())

	go watchStartedComplete(project, scanner, completeChannel, ctx)
	go func() {
		select {
		case <-completeChannel:
			log.Printf("%v is completed to start\n", project)
			watcherChannel <- true
		case <-time.After(20 * time.Second):
			log.Printf("%v is failed to start\n", project)
			cancel()
			watcherChannel <- false
		}
	}()

	cmd.Start()
	pm[project] = cmd.Process.Pid

	if ret := <- watcherChannel; ret {
		return nil
	}

	return errors.New("failed to start the " + project)
}

func watchStartedComplete(project string, scanner *bufio.Scanner, completeChannel chan struct{}, ctx context.Context) {


	completedMessage, _ := config.CompletedMessage()
	completedMessage = strings.Replace(completedMessage, "{prj}", project, 1)
	for scanner.Scan() {
		select {
		default:
			line := scanner.Text()
			if strings.EqualFold(completedMessage, line) {
				log.Println("checkmate ", project)
				completeChannel <- struct{}{}
				break
			}
		case <-ctx.Done():  // if cancel() execute
			log.Println("cancel to watch ",project," for go routine")
			return
		}
	}
}


func assertDependency(project string) error {
	dependencies, err := dependency.ReadDependency(project)
	if err != nil {
		return err
	}

	for _, dep := range *dependencies {
		if Status(dep).Status == "Down" {
			return errors.New("dep illegal, dep map : " +strings.Join(*dependencies, "->"))
		}
	}
	return nil
}

func Stop(project string) error {
	pid, err := findPid(project)
	if err != nil {
		return err
	}

	log.Println("pid : ", strconv.Itoa(pid))

	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		log.Println(err)
		return errors.New("failed to find process, project : " + project + ", pid : " + strconv.Itoa(pid))
	}

	// Kill the process
	err = proc.Kill()
	if err != nil {
		log.Println(err)
		return errors.New("failed to kill process, project : " + project + ", pid : " + strconv.Itoa(pid))
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

func Status(project string) *marine.ProjectStatus {
	pid, err := findPid(project)
	if err != nil {
		return &marine.ProjectStatus{Project: project, Status: "Down"}
	}
	process, err := process.NewProcess( int32(pid))
	if err != nil {
		return &marine.ProjectStatus{Project: project, Status: "Down"}
	}
	createTime, _ := process.CreateTime()
	log.Println("create time : ", strconv.Itoa(int(createTime)), ", pid : ", strconv.Itoa(pid))
	return &marine.ProjectStatus{Project: project, Status: "Running", Uptime: uptimeShortString(createTime), Pid: int32(pid)}
}

func Summary() (*marine.StatusSummary, error) {
	projects, err := config.Projects()
	if err != nil {
		return nil, err
	}

	c := make(chan *marine.ProjectStatus)
	for _, project := range projects {
		go func(prj string) {
			c <- Status(prj)
		}(project)
	}

	var summary []*marine.ProjectStatus
	for i:=0; i<len(projects); i++ {
		summary = append(summary, <-c)
	}
	return &marine.StatusSummary{Projects: summary}, nil
}

func findPid(project string) (int, error) {
	pid, err := findPidByMap(project)
	if err != nil {
		pid, err = findPidByName(project)
		if err != nil { // process가 존재하지 않음
			log.Println(project, "가 존재하지 않음")
			return -1, errors.New("not running project")
		}
	}
	return pid, nil
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
		if err!=nil {
			continue
		}
		if isRunningByName(cmdLine, project) {
			return errors.New("already started.., please start after kill process, project : " + project + ", pid : " + strconv.Itoa(int(pr.Pid)))
		}
	}
	return nil
}

func isRunningByName(cmdLine string, project string) bool {
	allowCommand, err := config.AllowCommand()
	if err != nil {
		return false
	}
	command := strings.Split(cmdLine, " ")[0]
	return strings.Contains(command, allowCommand) && strings.Contains(cmdLine, project)
}

func findPidByName(project string) (int, error) {
	processes, _ := process.Processes()
	for _, pr := range processes {
		var cmdLine string
		var err error
		cmdLine, err = pr.Cmdline()
		if err != nil {
			continue
		}
		if isRunningByName(cmdLine, project) {
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
	workingDir, err := config.WorkingDirectory()
	log.Println("workingDir : ", workingDir)
	if err != nil {
		return "./"
	}
	return  workingDir+ "/" + project
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
func uptime(startTime time.Time) time.Duration {
	return time.Since(startTime)
}

func uptimeShortString(startTime int64) string {
	duration := uptime(time.Unix(startTime/1000, 0))
	shortDuration, _ := durafmt.ParseStringShort(duration.String())
	return shortDuration.String()
}