package f

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/aidos-dev/words_counter/pkg/sorter"
)

func Test_PrintTable(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	/*
		25 items will be passed to func and it should
		print out only 20
	*/
	outputSet := sorter.SliceOfOutputs{
		{Word: []byte("the"), Number: 4284},
		{Word: []byte("and"), Number: 2192},
		{Word: []byte("of"), Number: 2185},
		{Word: []byte("a"), Number: 1861},
		{Word: []byte("to"), Number: 1685},
		{Word: []byte("in"), Number: 1366},
		{Word: []byte("i"), Number: 1056},
		{Word: []byte("that"), Number: 1024},
		{Word: []byte("his"), Number: 889},
		{Word: []byte("it"), Number: 821},
		{Word: []byte("he"), Number: 783},
		{Word: []byte("but"), Number: 616},
		{Word: []byte("was"), Number: 603},
		{Word: []byte("with"), Number: 595},
		{Word: []byte("s"), Number: 577},
		{Word: []byte("is"), Number: 564},
		{Word: []byte("for"), Number: 551},
		{Word: []byte("all"), Number: 542},
		{Word: []byte("as"), Number: 541},
		{Word: []byte("at"), Number: 458},
		{Word: []byte("this"), Number: 452},
		{Word: []byte("not"), Number: 419},
		{Word: []byte("him"), Number: 410},
		{Word: []byte("be"), Number: 384},
		{Word: []byte("on"), Number: 366},
	}

	PrintTable(outputSet)

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test

	// Manually fill the expected output
	expectedOutput := "   4284 the\n" +
		"   2192 and\n" +
		"   2185 of\n" +
		"   1861 a\n" +
		"   1685 to\n" +
		"   1366 in\n" +
		"   1056 i\n" +
		"   1024 that\n" +
		"    889 his\n" +
		"    821 it\n" +
		"    783 he\n" +
		"    616 but\n" +
		"    603 was\n" +
		"    595 with\n" +
		"    577 s\n" +
		"    564 is\n" +
		"    551 for\n" +
		"    542 all\n" +
		"    541 as\n" +
		"    458 at\n"

	// if string(out) != expectedOutput {
	// 	t.Errorf("incorrect output:\nexpected:\n%s\ngot:\n%s", expectedOutput, string(out))
	// }

	if !strings.EqualFold(string(out), expectedOutput) {
		t.Errorf("incorrect output:\nexpected:\n%s\ngot:\n%s", expectedOutput, string(out))
	}
}
