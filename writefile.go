package reloaded

import (
	"io/ioutil"
	"strings"
)

func Writefile(newSlice []string, file_output string) {
	final := strings.Join(newSlice, " ")

	output := []byte(final)

	ioutil.WriteFile(file_output, output, 0644)
}
