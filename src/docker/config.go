package docker

import (
	"fmt"
	"goconf/conf"
	"os"
	"path/filepath"
)

func init() {
	filePath, err := searchConfigFile()
	if nil != err {
		panic(err.Error())
	}
	file, err = conf.ReadConfigFile(filePath)
	if nil != err {
		panic(err.Error())
	}

}

var (
	file *conf.ConfigFile
)

func searchConfigFile() (string, error) {
	path, err := os.Getwd()
	if nil != err {
		return "", fmt.Errorf("Search config file error: %v", err.Error())
	}

	isWalkRoot := false
	confPath := ""
	findConfig := fmt.Errorf("find config file")
	for {
		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if nil != err {
				return err
			}

			if info.IsDir() && path != path {
				return filepath.SkipDir
			}
			if info.Name() == "test-on-docker.conf" {
				confPath = path
				return findConfig
			}
			return nil
		})
		parentPath = filepath.Dir(path)
		if isWalkRoot {
			break
		}
		if findConfig == err {
			return confPath, nil
		}
		if nil != err {
			return "", err
		}

		// "." or "\"
		if len(path) == 1 {
			isWalkRoot = true
		}
	}
	return "", fmt.Errorf("No config file")
}

func getString(section, option string) string {
	ret, err := file.GetString(section, option)
	if nil != err {
		panic(err.Error())
	}

	if "" == ret {
		panic(fmt.Sprintf("config(%v.%v) required\n", section, option))
	}

	return ret
}

func getHostPath() string {
	return getString("path", "host")
}

func getDockerPath() string {
	return getString("path", "docker")
}
