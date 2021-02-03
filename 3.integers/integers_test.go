package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	get := Add(2, 2)
	expected := 4

	if get != expected {
		t.Errorf("expected %v, get %v", expected, get)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
