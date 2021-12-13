package reloaded

import "fmt"

func Checker(arguments []string) bool {
	if len(arguments) < 3 {
		fmt.Println("Missing args")
		return false
	} else if len(arguments) > 3 {
		fmt.Println("Too many arguments")
		return false
	}
	return true
}
