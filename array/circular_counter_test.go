package array

import (
	"testing"
	"fmt"
)

func TestJosepheusFunc(t *testing.T) {
	t.Log("Test Circular Counter algorithm")
	temp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print("OUTPUT: ")
	josepheus(temp, 3)
	fmt.Println()
}
