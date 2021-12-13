package reloaded

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Writefile(newSlice []string, file_output string) {
	final := strings.Join(newSlice, " ")

	output := []byte(final)

	f, err := os.Create(os.Args[2] + ".txt")
	if err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(f)
	w.WriteString(string(output))
	w.Flush()

}
