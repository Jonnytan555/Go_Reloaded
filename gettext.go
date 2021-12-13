package reloaded

import (
	"fmt"
	"os"
)

func GetText(file_input string) string {
	file, err := os.Open(file_input)
	if err != nil {
		fmt.Printf("The mistake is: %v\n", err.Error())
	}

	aFile, _ := file.Stat()

	size := aFile.Size()

	arr := make([]byte, size)

	file.Read(arr)

	defer file.Close()

	return string(arr)
}
