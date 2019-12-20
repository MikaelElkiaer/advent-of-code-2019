package shared

import (
	"io/ioutil"
	"strings"
)

func GetInput(path string) string {
	data, _ := ioutil.ReadFile(path)

	return string(data)
}

func GetInputLines(path string) []string {
	data := GetInput(path)

	lines := strings.Split(data, "\n")

	return lines
}
