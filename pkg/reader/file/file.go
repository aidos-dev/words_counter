package file

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadFile() ([]byte, error) {

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	data, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	return data, err
}
