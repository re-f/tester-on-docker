package docker

import (
	"fmt"
	"goconf/conf"
	"os"
	"path/filepath"

	"github.com/ungerik/go-dry/dry"
)

var (
	config *conf.ConfigFile
)

func loadConfig() error {
	filePath, err := searchConfigFile()
	if nil != err {
		return fmt.Errorf("Load config error: %v", err.Error())
	}

	config, err = conf.ReadConfigFile(filePath)
	if nil != err {
		return fmt.Errorf("Load config error: %v", err.Error())
	}
	debugLog("config at %v", filePath)
	return nil
}

func searchConfigFile() (string, error) {
	curPath, err := os.Getwd()
	if nil != err {
		return "", fmt.Errorf("Search config file error: %v", err.Error())
	}
	for {
		confPath := filepath.Join(curPath, "test-on-docker.conf")
		if dry.FileExists(confPath) {
			return confPath, nil
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
	ret, err := config.GetString(section, option)
	if nil != err {
		panic(err.Error())
	}

	if "" == ret {
		panic(fmt.Sprintf("config(%v.%v) required\n", section, option))
	}

	return ret
}

func getHostPath() string {
	return filepath.ToSlash(getString("path", "host"))
}

func getBoot2DockerPath() string {
	return filepath.ToSlash(getString("path", "docker"))
}

func isDebug() bool {
	isDebug, _ := config.GetBool("global", "debug")
	return isDebug
}
func isRebuild() bool {
	rebuild, _ := config.GetBool("image", "rebuild")
	return rebuild
}
func getSections() []string {
	return config.GetSections()

}

func getImage() image {
	ins := image{}
	ins.name = getString("image", "name")
	ins.os = getString("image", "os")
	ins.arch = getString("image", "arch")
	return ins
}
