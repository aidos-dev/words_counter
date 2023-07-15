package app

import (
	"fmt"

	"github.com/aidos-dev/words_counter/internal/models"
	"github.com/aidos-dev/words_counter/pkg/counter"
	"github.com/aidos-dev/words_counter/pkg/printer/f"
	"github.com/aidos-dev/words_counter/pkg/reader/file"
	"github.com/aidos-dev/words_counter/pkg/reader/word"
)

func Run() {

	const op = "Run"

	data := []byte{49, 50, 51, 52, 53}

	f.Println(data)

	text, _ := file.ReadFile()

	wordCh := make(chan models.WordBytes)
	outputCh := make(chan []models.Output)

	go word.CleanWords(text, wordCh)

	go counter.Count(wordCh, outputCh)

	outputSet := <-outputCh

	fmt.Printf("ouputSet value: %v\n", outputSet)

	for _, el := range outputSet {

		fmt.Print(string(el.Word))
		fmt.Print(": ")
		fmt.Println(el.Number)

		f.Println(el.Word)
		num := f.IntToBytes(el.Number)
		f.Println(num)
	}

	fmt.Printf("%s: Done\n", op)

}
