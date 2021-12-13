package reloaded

func ToSlice(text string) []string {
	var newRune []rune
	var theString []string

	for _, char := range text {
		if char == ' ' {
			if newRune != nil {
				theString = append(theString, string(newRune))
				newRune = nil
			}
		} else {
			newRune = append(newRune, char)
		}
	}

	if newRune != nil {
		theString = append(theString, string(newRune))
	}

	var fixedSlice []string

	for _, word := range theString {
		if word[len([]rune(word))-1] == ')' && (word[0] >= '0' && word[0] <= '9') {
			fixedSlice[len(fixedSlice)-1] += " " + word
		} else {
			fixedSlice = append(fixedSlice, word)
		}
	}

	return fixedSlice
}
