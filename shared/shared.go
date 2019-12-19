package shared

import (
	"io/ioutil"
	"strings"
)

func GetInput(path string) []string {
	data, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(data), "\n")

	return lines
}
