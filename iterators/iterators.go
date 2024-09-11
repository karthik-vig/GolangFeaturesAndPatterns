package iterators

import (
	"fmt"
	"iter" // for iterator data type signature
	"slices"
)

// basic iterator
func basicOneItemIteratorFunc(yield func(byte) bool) {
	someSlice := []byte{1, 2, 3, 4}
	for _, val := range someSlice {
		if !yield(val) {
			// do clean up if needed
			return
		}
	}
}

// function that returns an iterator
func funcThatReturnsIterator(counter *int) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		someSlice := []byte{1, 2, 3, 4}
		for _, val := range someSlice {
			(*counter)++
			if !yield(val) {
				// do clean up if needed
				return
			}
		}
	}
}

// iterator that returns two values
func funcThatReturnsIteratorWith2Value() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		someSlice := []int{1, 2, 3, 4}
		for idx, val := range someSlice {
			if !yield(idx+1, val*10) {
				// clean up code
				return
			}
		}
	}
}

// iterators with second value for error reporting
type iteratorError struct{}

func (err *iteratorError) Error() string {
	return "Iterator Error!!!"
}

func funcThatReturnsIteratorAndError() iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		someSlice := []int{1, 2, 3, 4}
		for _, val := range someSlice {
			err := error(nil)
			if val == 3 {
				err = &iteratorError{}
			}
			if !yield(val*20, err) {
				// clean up code
				return
			}
		}
	}
}

// using in-built slice iterators
func sliceIterator() {
	fmt.Println("Iterating through idx and value")
	someSlice := []int{1, 2, 3}
	for idx, val := range slices.All(someSlice) {
		fmt.Printf("The index is: %d and the value is: %d\n", idx, val)
	}
	fmt.Println("Iterating through value alone")
	for val := range slices.Values(someSlice) {
		fmt.Println(val)
	}
}

func IteratorExample() {
	fmt.Println("This is an example of one item yielding iterator")
	for oneItemVal := range basicOneItemIteratorFunc {
		fmt.Println(oneItemVal)
	}

	fmt.Println("Example of an function that returns an iterator")
	var counter int
	for oneItemVal := range funcThatReturnsIterator(&counter) {
		fmt.Println(oneItemVal)
	}
	fmt.Println("The counter value is: ", counter)

	fmt.Println("Example of a function that returns a iterator that emits two values")
	for idx, val := range funcThatReturnsIteratorWith2Value() {
		fmt.Printf("The index is %d and value is %d\n", idx, val)
	}

	fmt.Println("Example of a function that returns a iterator with second value as a error variable")
	for val, err := range funcThatReturnsIteratorAndError() {
		if err != nil {
			fmt.Println("Error has occurred, the error is:\n", err.Error())
			break
		}
		fmt.Println(val)
	}

	fmt.Println("Example of using slice iterator:")
	sliceIterator()
}
