package helper

import (
	"cdpui/config"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !IsDir(dir) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

func RemoveFile(file string) error {
	if FileIsExisted(file) {
		return os.Remove(file)
	}
	return nil
}

func RemoveDir(dir string) error {
	if IsDir(dir) {
		return os.RemoveAll(dir)
	}
	return nil
}

func GetWorkDir(index int) string {
	return fmt.Sprintf("%s/%v_task", config.ResourcesDir, index)
}

func OpenFileManager(index int) error {
	//TODO  mac linux OpenFileManager Implement
	dir := strings.ReplaceAll(GetWorkDir(index),"/","\\")
	cmd := exec.Command(`explorer.exe`, `select,`, dir)
	return cmd.Run()
}
