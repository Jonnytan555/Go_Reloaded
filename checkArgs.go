package reloaded

import "fmt"

func Checker(arguments []string) bool {
	if len(arguments) < 3 {
		fmt.Println("Missing args")
		fmt.Println("Usage: go run . [INPUT FILE] [OUTPUT FILE]")
		fmt.Println("EX: go run . standard output")
		return false
	} else if len(arguments) > 3 {
		fmt.Println("Too many arguments")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . standard output")
		return false
	} else {
		return true
	}
}
