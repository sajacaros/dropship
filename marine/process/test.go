package process

import (
	"github.com/shirou/gopsutil/process"
	"log"
	"strconv"
)

func test() {
	process, _ := process.NewProcess( 7055)
	createTime, _ := process.CreateTime()
	pid := process.Pid
	log.Println("create time : ", strconv.Itoa(int(createTime)), ", pid : ", strconv.Itoa(int(pid)))
}

