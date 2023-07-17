package word

import (
	"reflect"
	"testing"

	"github.com/aidos-dev/words_counter/internal/models"
)

func Test_CleanWords(t *testing.T) {
	testCases := []struct {
		name          string
		dirtyText     []byte
		expectedWords []models.WordBytes
	}{
		{
			name:          "unprintable",
			dirtyText:     []byte(" helLo  WoRlD   "),
			expectedWords: []models.WordBytes{[]byte("hello"), []byte("world")},
		},
		{
			name:          "symbols",
			dirtyText:     []byte("! @ # $hello %^& *( ) WOrLd- _ + = "),
			expectedWords: []models.WordBytes{[]byte("hello"), []byte("world")},
		},
		{
			name:          "numbers",
			dirtyText:     []byte("123456789Hello 123456789 world 1234567890"),
			expectedWords: []models.WordBytes{[]byte("hello"), []byte("world")},
		},
	}

	wordCh := make(chan models.WordBytes)

	receivedWords := []models.WordBytes{}

	for _, tc := range testCases {

		go CleanWords(tc.dirtyText, wordCh)

		for word := range wordCh {
			receivedWords = append(receivedWords, word)
		}

		if !reflect.DeepEqual(tc.expectedWords, receivedWords) {
			t.Errorf("Test case: '%s' failed: Expected %v, but got %v", tc.name, tc.expectedWords, receivedWords)
		}
	}
}
