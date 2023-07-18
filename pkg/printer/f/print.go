package f

import (
	"os"

	"github.com/aidos-dev/words_counter/pkg/sorter"
)

func PrintTable(outputSet sorter.SliceOfOutputs) {
	maxWordWidth := len(IntToBytes(outputSet[0].Number))
	oneSpace := []byte{' '}

	for i := 0; i < len(outputSet); i++ {
		if i == 20 {
			break
		}

		numPrintable := IntToBytes(outputSet[i].Number)

		Print(addSpaces(numPrintable, maxWordWidth))
		Print(numPrintable)
		Print(oneSpace)
		Println(outputSet[i].Word)
	}
}

func addSpaces(word []byte, maxWordWidth int) []byte {
	/*
		spaces have 3 spaces from the begining so even
		the largest number will have at least 3 spaces at the beginning
	*/
	spaces := []byte{' ', ' ', ' '}

	moreSpaces := maxWordWidth - len(word)

	for i := 0; i < moreSpaces; i++ {
		spaces = append(spaces, ' ')
	}

	return spaces
}

func Print(sliceOfBytes []byte) {
	os.Stdout.Write(sliceOfBytes)
}

func Println(sliceOfBytes []byte) {
	sliceOfBytes = append(sliceOfBytes, '\n')

	os.Stdout.Write(sliceOfBytes)
}

func IntToBytes(n int) []byte {
	if n == 0 {
		return []byte{'0'}
	}

	var bytes []byte
	negative := false

	if n < 0 {
		negative = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		byteDigit := '0' + byte(digit)
		bytes = append([]byte{byteDigit}, bytes...)
		n /= 10
	}

	if negative {
		bytes = append([]byte{'-'}, bytes...)
	}

	return bytes
}

func PrintFileErr() {
	fileErr := []byte{'e', 'r', 'r', 'o', 'r', ':', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', 't', 'o', ' ', 'a', ' ', 'r', 'e', 'a', 'd', ' ', 'f', 'i', 'l', 'e'}

	Println(fileErr)
}
