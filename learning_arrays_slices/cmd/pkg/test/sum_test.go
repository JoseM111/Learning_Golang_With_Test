package test

/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */
import (
	. "fmt"
	"reflect"
	"testing"
)

func TestSumArray(test *testing.T) {
	// ™ ____________
	/* Arrays have a fixed capacity which you define when you
	   declare the variable. We can initialize an array in two ways:
	
		• [N]type{value1, value2, ..., valueN} e.g.
		• numbersArray := [5]int{1, 2, 3, 4, 5}
		• [...]type{value1, value2, ..., valueN} e.g.
		• numbersArray := [...]int{1, 2, 3, 4, 5}
	*/
	numbersArray := [5]int{1, 2, 3, 4, 5}
	
	got := SumArray(numbersArray)
	want := 15
	
	if got != want {
		// It is sometimes useful to also print the inputs to the function
		// in the error message. Here, we are using the %v placeholder to
		// print the "default" format, which works well for arrays
		test.Errorf("got %d want %d given, %v", got, want, numbersArray)
	}
	
	Printf("want: %d\ngot: %v\n", want, got)
}

// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func TestSumSlice(test *testing.T) {
	// ™ ____________
	test.Run("slice collection of size 5", func(test *testing.T) {
		// ™ ____________
		numbersSlice := []int{1, 2, 3, 4, 5}
		
		got := SumSlice(numbersSlice)
		want := 15
		
		if got != want {
			test.Errorf("got %d want %d given, %v", got, want, numbersSlice)
		}
		
		Printf("want: %d\ngot: %v\n", want, got)
	})
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func TestSumAll(test *testing.T) {
	// ™ ____________
	sumSlice1 := []int{1, 2}
	sumSlice2 := []int{0, 9}
	
	got := SumAllSlice(sumSlice1, sumSlice2)
	want := []int{3, 9}
	
	if !reflect.DeepEqual(got, want) {
		test.Errorf("got %v want %v", got, want)
	}
	
	Printf("want: %d\ngot: %v\n", want, got)
}

// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

/*
* It will calculate the totals of the "tails" of each slice.
* The tail of a collection is all items in the collection except
* the first one */
func TestSumAllTailSlices(test *testing.T) {
	// ™ ____________
	checkSum := func(test *testing.T, got, want []int) {
		/* Helper marks the calling function as a test helper function.
		 * When printing file and line information, that function will
		 * be skipped. Helper may be called simultaneously from multiple goroutines.*/
        test.Helper()
        if !reflect.DeepEqual(got, want) {
            test.Errorf("got %v want %v", got, want)
        }
        
        Printf("want: %d\ngot: %v\n", want, got)
    }
	
	sumSlice1 := []int{1, 2}
	sumSlice2 := []int{0, 9}
	
	got := SumAllTailSlices(sumSlice1, sumSlice2)
	want := []int{2, 9}
	
	checkSum(test, got, want)
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func BenchmarkRepeat(benchmark *testing.B) {

}

/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */
