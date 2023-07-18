package app

import (
	"os"
	"sort"

	"github.com/aidos-dev/words_counter/internal/models"
	"github.com/aidos-dev/words_counter/pkg/counter"
	"github.com/aidos-dev/words_counter/pkg/printer/f"
	"github.com/aidos-dev/words_counter/pkg/reader/file"
	"github.com/aidos-dev/words_counter/pkg/reader/word"
	"github.com/aidos-dev/words_counter/pkg/sorter"
)

func Run() {
	if len(os.Args) != 2 {
		return
	}

	text, err := file.ReadFile()
	if err != nil {
		f.PrintFileErr()
		return
	}

	wordCh := make(chan models.WordBytes)
	outputCh := make(chan []models.Output)

	go word.CleanWords(text, wordCh)

	go counter.Count(wordCh, outputCh)

	var outputSet sorter.SliceOfOutputs

	outputSet = <-outputCh

	sort.Sort(outputSet)

	f.PrintTable(outputSet)
}
