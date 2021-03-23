package config

import "os"

var ResourcesDir string
var TaskFile string

func init() {

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ResourcesDir = currentDir + "/resources"
	TaskFile = currentDir + "/Task.json"
}
