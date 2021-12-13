package reloaded

import "strings"

func Modification(modifier string, newSlice []string, multiplier int) {
	switch modifier {

	case "cap":
		for k := 1; k <= multiplier; k++ {
			newSlice[len(newSlice)-k] = Capitalize(newSlice[len(newSlice)-k])
		}

	case "up":
		for k := 1; k <= multiplier; k++ {
			newSlice[len(newSlice)-k] = strings.ToUpper(newSlice[len(newSlice)-k])
		}
	case "low":
		for k := 1; k <= multiplier; k++ {
			newSlice[len(newSlice)-k] = strings.ToLower(newSlice[len(newSlice)-k])
		}
	case "hex":
		for k := 1; k <= multiplier; k++ {
			newSlice[len(newSlice)-1] = Hex(newSlice[len(newSlice)-1])
		}
	case "bin":
		for k := 1; k <= multiplier; k++ {
			newSlice[len(newSlice)-1] = Bin(newSlice[len(newSlice)-1])
		}
	}
}
