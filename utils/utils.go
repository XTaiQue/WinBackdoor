package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	ps "github.com/mitchellh/go-ps"

	"github.com/shirou/gopsutil/host"
)

func System(command string) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		fmt.Println("[!] Windows Required !")
	}
}

func ClearConsole() {
	if runtime.GOOS == "windows" {
		System("cls")
	} else {
		fmt.Println("[!] Windows Required !")
	}
}

func Pwd() (string, error) {
	getcwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return getcwd, nil
}

func CD(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	os.Chdir(path)
	return true
}

func GetUser() (string, error) {
	user, err := user.Current()

	if err != nil {
		return "", nil
	}

	if runtime.GOOS == "windows" {
		split := strings.Split(user.Username, "\\")
		return split[1], nil
	} else {
		fmt.Println("Windows Required !")
		return "", nil
	}
}

func GetHostname() (string, error) {
	hostStat, err := host.Info()

	if err != nil {
		return "", nil
	}

	return hostStat.Hostname, nil
}

func GetPlatformVersion() (string, error) {
	hostStat, err := host.Info()

	if err != nil {
		return "", nil
	}

	return hostStat.PlatformVersion, nil
}

func GetPlatform() (string, error) {
	hostStat, err := host.Info()

	if err != nil {
		return "", nil
	}

	return hostStat.Platform, nil
}

func GetPlatformFamily() (string, error) {
	hostStat, err := host.Info()

	if err != nil {
		return "", nil
	}

	return hostStat.PlatformFamily, nil
}

func Ls(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func Cat(path string) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}

func Mkdir(dir string) error {

	err := os.MkdirAll(dir, os.ModeDir)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func Rm(file string) error {
	err := os.Remove(file)

	if err != nil {
		return err
	}

	return nil
}

func RmDir(path string) error {
	err := os.RemoveAll(path)

	if err != nil {
		return err
	}

	return nil
}

func Rename(file string, newfile string) error {
	err := os.Rename(file, newfile)
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(file string, data []byte) error {
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		return err
	}

	write := ioutil.WriteFile(file, data, 0644)

	if write != nil {
		return write
	}
	return nil
}

func RmA(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetLocalAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		return "", err
	}

	ip := conn.LocalAddr().String()
	return ip, nil
}

func ShowPs() {
	processList, err := ps.Processes()

	if err != nil {
		fmt.Println(err)
	}

	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("\t%d\t%s\n", process.Pid(), process.Executable())
	}
}

func GetProcess() ([]ps.Process, error) {
	processList, err := ps.Processes()

	if err != nil {
		fmt.Println(err)
	}

	return processList, nil
}

func FindProc(pid int) (int, int, string, error) {
	pe, e := ps.FindProcess(pid)

	if e != nil {
		return 0, 0, "", nil
	}

	return pe.PPid(), pe.Pid(), pe.Executable(), nil
}
