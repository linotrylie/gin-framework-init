package sliceutil

import (
	"fmt"
	"testing"
)

func TestSliceFinder(t *testing.T) {
	testData := []string{
		"d",
		"d",
	}

	duplicated := SliceEleIsDuplicated(testData)
	fmt.Println(duplicated)
}
