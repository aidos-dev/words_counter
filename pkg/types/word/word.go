package word

import "os"

type WordStr struct {
	Word []byte
}

func (w *WordStr) Print() {
	os.Stdout.Write(w.Word)
}
