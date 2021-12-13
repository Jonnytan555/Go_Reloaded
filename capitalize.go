package reloaded

func Capitalize(s string) string {
	sen := []rune(s)

	// all capitalised letters are lower cased
	for i, char := range sen {
		if char >= 'A' && char <= 'Z' {
			sen[i] = char + 32
		}
	}

	for i, char := range sen {

		// if the first character is lower cased
		// capitalise the value
		if sen[0] >= 'a' && sen[0] <= 'z' {
			sen[0] = sen[0] - 32
		}

		// if the the value is not alphanumeric
		if (char < '0') || (char > '9' && char < 'A') || (char > 'Z' && char < 'a') || (char > 'z') {
			// if the iteration is at the final position, break
			if i == len(sen)-1 {
				break
			}
			// if the next letter after the iteration is a lower number, capiatlise it
			if sen[i+1] >= 'a' && sen[i+1] <= 'z' {
				sen[i+1] = sen[i+1] - 32
			}
		}
	}
	// return a string
	return string(sen)
}
