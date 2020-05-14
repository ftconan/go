// Author: magician
// File:   sort_test.go
// Date:   2020/5/14
package treesort_test

import (
	"golang/go_program/ch4/treesort"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
