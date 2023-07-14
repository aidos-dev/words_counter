package app

import (
	"github.com/aidos-dev/words_counter/pkg/printer/f"
	"github.com/aidos-dev/words_counter/pkg/reader/file"
	"github.com/aidos-dev/words_counter/pkg/reader/word"
)

func Run() {
	data := []byte{49, 50, 51, 52, 53}

	f.Println(data)

	text, _ := file.ReadFile()

	testWords := word.CleanWords(text)

	for _, el := range testWords {
		f.Println(el)
	}

}
