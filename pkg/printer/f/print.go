package f

import "os"

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
