package counter

import (
	"bytes"

	"github.com/aidos-dev/words_counter/internal/models"
)

func Count(wordCh chan models.WordBytes, outputCh chan []models.Output) {
	const op = "Count"

	outputSet := []models.Output{}

	for {
		value, ok := <-wordCh

		if !ok {
			// Channel is closed, no more values will be received

			break
		}

		countFrequency(&outputSet, value)

	}

	outputCh <- outputSet
}

/*
countFrequency increments number of word occurence
*/
func countFrequency(outputSet *[]models.Output, value models.WordBytes) {
	const op = "Count Frequency"

	for i := range *outputSet {
		if bytes.Equal((*outputSet)[i].Word, value) {
			(*outputSet)[i].Number++
			return
		}
	}

	var tempValue models.Output
	tempValue.Word = value
	tempValue.Number = 1

	*outputSet = append(*outputSet, tempValue)
}
