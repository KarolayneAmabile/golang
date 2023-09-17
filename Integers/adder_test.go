package integers // group functions for working with integers such as add

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum) // using %d because we want it to print a string
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// GO examples are executed just like tests so you can e confident examples reflect what the
// code actually does. Examples are compiled (and optionally executed) as part of a package's test suite
// If your code changes so that example is no longer valid, your build will fail.
// The exemple function will be not executed if you remove the comment "// Output: 6"
// By adding this code the example will appear in the documentation inside godoc, making your code
// even more accessible.
