// +build !windows

package docker

func executeOnDocker(str string) (string, error) {
	return execute(str)
}

func getLineEnd() string {
	return "\n"
}
func getScriptSuffix() string {
	return "sh"
}
