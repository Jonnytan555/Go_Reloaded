package reloaded

import (
	"fmt"
	"strconv"
)

func Bin(newSlice string) string {
	output, err := strconv.ParseInt(newSlice, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(output))
}
