package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

func main() {
	var hexaNumber string
	fmt.Print("Enter Hexadecimal Number:")
	fmt.Scanln(&hexaNumber)

	output, err := strconv.ParseInt(hexaNumberToInteger(hexaNumber), 16, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Output %d", output)
}
