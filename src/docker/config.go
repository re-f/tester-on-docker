package docker

import (
	"fmt"
	"goconf/conf"
	"os"
	"path/filepath"
)

var (
	file *conf.ConfigFile
)

func searchConfigFile() (string, error) {
	curPath, err := os.Getwd()
	if nil != err {
		return "", fmt.Errorf("Search config file error: %v", err.Error())
	}

	confPath := ""
	findConfigFlag := fmt.Errorf("find config file")
	for {
		// @todo don't walk
		err := filepath.Walk(curPath, func(path string, info os.FileInfo, err error) error {
			if nil != err {
				return err
			}
			if info.IsDir() && path != curPath {
				return filepath.SkipDir
			}
			if info.Name() == "test-on-docker.conf" {
				confPath = filepath.Join(curPath, info.Name())
				return findConfigFlag
			}
			return nil
		})
		if findConfigFlag == err {
			return confPath, nil
		}
		if nil != err {
			return "", fmt.Errorf("Search config file error", err.Error())
		}
		parentPath := filepath.Dir(curPath)
		if parentPath == curPath {
			break
		}
		curPath = parentPath
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

func isDebug() bool {
	isDebug, _ := file.GetBool("global", "debug")
	return isDebug
}
func getSections() []string {
	return file.GetSections()

}

func getImage() image {
	ins := image{}
	ins.name = getString("image", "name")
	ins.os = getString("image", "os")
	ins.arch = getString("image", "arch")
	return ins
}
