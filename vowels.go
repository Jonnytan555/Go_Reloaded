package reloaded

import "fmt"

func Vowels(puncEdit []string) []string {
	for j := 0; j < len(puncEdit); j++ {
		if puncEdit[j] == "a" && IsVowel((puncEdit[j+1])[0]) {
			puncEdit[j] = "an"
		}
	}
	fmt.Println(puncEdit)
	return puncEdit
}
