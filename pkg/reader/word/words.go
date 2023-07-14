package word

import "github.com/aidos-dev/words_counter/internal/models"

/*
CleanWords accepts row text as a slice of bytes
and returns a slice of "Words" without any unprintable
characters
*/
func CleanWords(dirtyText []byte) []models.Word {
	var cleanWords []models.Word

	var tempFiller []byte

	for i := 0; i < len(dirtyText); i++ {
		if (dirtyText[i] >= 'A' && dirtyText[i] <= 'Z') || (dirtyText[i] >= 'a' && dirtyText[i] <= 'z') {
			tempFiller = append(tempFiller, toLow(dirtyText[i]))
		} else {
			if len(tempFiller) > 0 {
				cleanWords = append(cleanWords, tempFiller)
				tempFiller = nil
			}

		}
	}

	return cleanWords
}

/*
Calling func toLow for every letter helps to make all words in
lower case, to make words counting functions case insensitive
*/
func toLow(symb byte) byte {
	if symb >= 'A' && symb <= 'Z' {
		symb += 32
	}

	return symb
}
