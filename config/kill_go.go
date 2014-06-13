package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func killProcessByPid(pid int) {
	pidStr := fmt.Sprintf("%d", pid)
	cmd := exec.Command("kill", pidStr)
	output, _ := cmd.Output()
	ret := string(output)
	if len(ret) > 1 {
		fmt.Println(ret)
	}
}

func findProcessByName(name string) (pid int) {
	cmdStr := fmt.Sprintf("ps -A | grep -m1 %s | grep -v grep | grep -v %s | awk '{print $1}'", name, os.Args[0])
	cmd := exec.Command("sh", "-c", cmdStr)
	if output, err := cmd.Output(); err != nil {
		fmt.Println(err)
	} else {
		ret := string(output)
		ret = strings.Replace(ret, "\n", "", -1)
		if _pid, _err := strconv.ParseInt(ret, 10, 64); _err != nil {
			// fmt.Println(_err)
		} else {
			pid = int(_pid)
		}
	}
	return pid
}

func killGoProcess(names []string) {
	for _, value := range names {
		pid := findProcessByName(value)
		if pid > 0 {
			killProcessByPid(pid)
			fmt.Printf("%s is killed\n", value)
		} else {
			fmt.Printf("%s is not launched\n", value)
		}
	}
}

func main() {
	// fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println("args invalid")
	} else {
		killGoProcess(os.Args[1:])
	}
}
