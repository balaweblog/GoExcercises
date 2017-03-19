package math

import (
	"fmt"
	"testing/quick"
    "testing"
)

/*
//simple test
//go test math
func TestAdd(t *testing.T) {
	result := Add(1, 3, 4, 4)
	if result != 12 {
		t.Fatal("error happened ")
	}
}
// benchmark test
//go test -bench=.
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 7)
	}
    fmt.Println("bye")

}
//example test
//go test
func ExampleAdd() {
	fmt.Println(Add(4, 7))
	// Output:
	// 11
}
*/
// blackbox testing testing/quick package needed
func TestAddBlackbox(t *testing.T) {
    fmt.Println("entered")
    err := quick.Check(a, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func a(x, y int) bool {
	return Add(x, y) == x+y
}