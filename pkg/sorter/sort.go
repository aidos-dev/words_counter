package sorter

import "github.com/aidos-dev/words_counter/internal/models"

/*
SliceOfOutputs struct satisfies the sort.Interface interface.
The Len(), Swap(), and Less() methods are required to implement the sort.Interface interface.
Implementing the sort.Interface interface enables you to use the sort.Sort() function from the sort package,
which internally uses an efficient sorting algorithm like Quicksort or Heap Sort to sort the elements in the slice
according to your specified order.

By implementing these methods, you define the behavior of the sorting algorithm for your specific type.
In the case of "SliceOfOutputs", the Less() method compares the Number field of Output structs,
allowing you to sort them in descending order based on that field.
*/

type SliceOfOutputs []models.Output

func (a SliceOfOutputs) Len() int {
	return len(a)
}

func (a SliceOfOutputs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a SliceOfOutputs) Less(i, j int) bool {
	return a[i].Number > a[j].Number
}
