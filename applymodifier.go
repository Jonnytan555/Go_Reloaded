package reloaded

import (
	"strconv"
)

func ApplyModifiers(aString []string) []string {
	var isModifier bool
	var multiplier int
	var modifier string
	var newSlice []string

	for _, word := range aString {
		modifier = ""

		multiplier = 1

		isModifier = false

		if word[0] == '(' && word[len(word)-3] == ' ' {
			for j, char := range word {
				if char == ',' {
					modifier = word[1:j]
				}
			}
			isModifier = true
			multiplier, _ = strconv.Atoi(string(word[len(word)-2]))
		} else if word[0] == '(' {
			modifier = word[1 : len(word)-1]
			isModifier = true
		}

		if !isModifier {
			newSlice = append(newSlice, word)
		} else {
			Modification(modifier, newSlice, multiplier)
		}
	}

	return newSlice
}
