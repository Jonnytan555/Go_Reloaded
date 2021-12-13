package reloaded

func Punctuation(newSlice []string) []string {
	// Initialise an empty slice of strings
	var puncEdit []string

	// Initialise found as false and count to zero
	var found bool
	count := 0

	// Iterate through the words in the slice of strings
	for _, word := range newSlice {
		found = false
		for _, char := range word {
			// if the character in the word is a piece of puntuation, initialise punCount to 0
			if (char == ',' || char == ':' || char == '.' || char == ';' || char == '!' || char == '?') && !found {
				punCount := 0
				// Iterate through the rest of the words in the string, increase the puncCount
				for _, punc := range word {
					if punc == ',' || punc == ':' || punc == '.' || punc == ';' || punc == '!' || punc == '?' {
						punCount++
					}
				}

				// If the length of the word is equal to the puntuation counter,
				// Concatinate the previous word before and the previous list of punctuation
				if punCount == len(word) {
					puncEdit[count-1] = puncEdit[count-1] + string(word)
				} else {
					// if the length of the word is not the same as the punc count
					// concatinate the character to the end of the last word
					puncEdit[count-1] = puncEdit[count-1] + string(char)

					// if the words length is greater than 1 
					// append the word (slicing off the piece of punctuation) to puncEdit
					if len(word) > 1 {
						puncEdit = append(puncEdit, word[1:])
						count++
					}
				}
				found = true

			// If puntuation is not found, append the word to puncEdit slice of strings 
			} else if !found {
				puncEdit = append(puncEdit, word)
				found = true
				count++
			}
		}
	}
	return puncEdit
}
