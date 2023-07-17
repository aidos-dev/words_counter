package counter

import (
	"reflect"
	"testing"

	"github.com/aidos-dev/words_counter/internal/models"
)

func Test_Count(t *testing.T) {
	testCases := []struct {
		name              string
		inputWords        []models.WordBytes
		expectedOutputSet []models.Output
	}{
		{
			name:              "hello",
			inputWords:        []models.WordBytes{[]byte("hello"), []byte("world"), []byte("world")},
			expectedOutputSet: []models.Output{{Word: []byte("hello"), Number: 1}, {Word: []byte("world"), Number: 2}},
		},
		{
			name:              "hello 2",
			inputWords:        []models.WordBytes{[]byte("hello"), []byte("world"), []byte("world"), []byte("there"), []byte("there"), []byte("there"), []byte("there")},
			expectedOutputSet: []models.Output{{Word: []byte("hello"), Number: 1}, {Word: []byte("world"), Number: 2}, {Word: []byte("there"), Number: 4}},
		},
	}

	// var outputSet sorter.SliceOfOutputs

	for _, tc := range testCases {

		wordCh := make(chan models.WordBytes)
		outputCh := make(chan []models.Output)

		go func() {
			for _, el := range tc.inputWords {
				wordCh <- el
			}

			close(wordCh)
		}()

		go Count(wordCh, outputCh)

		outputSet := <-outputCh

		if !reflect.DeepEqual(tc.expectedOutputSet, outputSet) {
			t.Errorf("Test case: '%s' failed: Expected %v, but got %v", tc.name, tc.expectedOutputSet, outputSet)
		}
	}
}
