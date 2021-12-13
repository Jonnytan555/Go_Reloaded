package reloaded

import (
	"os"
)

func GoReloaded() {
	arguments := os.Args

	// Check that the input arguments are of length 3 (name of file,
	//	input file name and output file name)

	if !Checker(arguments) {
		return
	}

	file_input := os.Args[1]
	file_output := os.Args[2]

	text := GetText(file_input)

	aString := ToSlice(text)

	var newSlice []string
	var puncEdit []string

	newSlice = ApplyModifiers(aString)

	puncEdit = Punctuation(newSlice)
	puncEdit = Quotes(puncEdit)
	edit := Vowels(puncEdit)

	Writefile(edit, file_output)
}
