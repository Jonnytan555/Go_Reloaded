package reloaded

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func Quotes(puncEdit []string) []string {
	openQuote := true

	for i := 0; i < len(puncEdit); i++ {
		if puncEdit[i] == "'" {
			if openQuote {
				puncEdit[i+1] = "'" + puncEdit[i+1]
				puncEdit = remove(puncEdit, i)
				i--
				openQuote = false
			} else if !openQuote {
				puncEdit[i-1] = puncEdit[i-1] + "'"
				puncEdit = remove(puncEdit, i)
				i--
			}
		}
	}
	return puncEdit
}
