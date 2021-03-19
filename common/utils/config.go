package utils

import "runtime"

func GetConfigPath() string {
	path := "../config"
	if runtime.GOOS == "windows" {
		path = "./"
	}
	return path
}
