package strcount

import (
	"fmt"
	"testing"
)

func TestStrCount(t *testing.T) {
	count := StrCount("umbrella", "l")
	expected := 2
	
	if count != expected {
		t.Errorf("expecetd %q but got %q", expected, count)
	}
}

func ExampleStrCount(){
	count := StrCount("umbrella", "l")
	fmt.Println(count)
	// Output: 2
}

func BenchmarkStrCount(b *testing.B) {
	for i:=0; i<b.N; i++{
		StrCount("umbrella", "l")
	} 
}