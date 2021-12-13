package reloaded

import (
	"fmt"
	"strconv"
)

func Hex(newSlice string) string {
	output, err := strconv.ParseInt(newSlice, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(output))
}
