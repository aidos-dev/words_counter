package f

import "os"

func Print(sliceOfBytes []byte) {
	os.Stdout.Write(sliceOfBytes)
}

func Println(sliceOfBytes []byte) {

	sliceOfBytes = append(sliceOfBytes, '\n')

	os.Stdout.Write(sliceOfBytes)
}
