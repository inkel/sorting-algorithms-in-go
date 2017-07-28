package sort

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"
)

var expected, seed []int

var suite = map[string]func([]int){
	"Bubble":    Bubble,
	"Selection": Selection,
	"Insertion": Insertion,
	"Shell":     Shell,
	"Comb":      Comb,
	"Merge":     testMergeSort,
	"Go":        sort.Ints,
}

func init() {
	rand.Seed(time.Now().UnixNano())

	if os.Getenv("SIZE") == "" {
		seed = []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	} else {
		size, err := strconv.Atoi(os.Getenv("SIZE"))
		if err != nil {
			panic(err)
		}

		seed = make([]int, size)

		for i := 0; i < size; i++ {
			seed[i] = rand.Int()
		}
	}

	fmt.Printf("SIZE: %d\n", len(seed))

	expected = make([]int, len(seed))
	copy(expected, seed)
	sort.Ints(expected)
}

func getItems() []int {
	items := make([]int, len(seed))
	copy(items, seed)
	return items
}

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func testMergeSort(items []int) {
	for i, v := range Merge(items) {
		items[i] = v
	}
}

func TestAlgorithms(t *testing.T) {
	for n, fn := range suite {
		t.Run(n, func(t *testing.T) {
			items := getItems()
			fn(items)
			if !isSorted(items) {
				t.Errorf("\nexp: %v\ngot: %v", expected, items)
			}
		})
	}
}

func BenchmarkAlgorithms(b *testing.B) {
	seed := getItems()
	items := make([]int, len(seed))

	b.ResetTimer()

	for n, fn := range suite {
		b.Run(n, func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				copy(items, seed)
				fn(items)
			}
		})
	}
}

func BenchmarkCopy(b *testing.B) {
	seed := getItems()
	items := make([]int, len(seed))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		copy(items, seed)
	}
}

// func TestAlgorithms_Quick(t *testing.T) {
// 	for n, fn := range suite {
// 		t.Run(n, func(t *testing.T) {
// 			err := quick.Check(func(items []int) bool {
// 				fn(items)
// 				return isSorted(items)
// 			}, nil)
// 			if err != nil {
// 				t.Error(err)
// 				// e := err.(*quick.CheckError)
// 				// items := make([]int, len(e.In))
// 				// for i, v := range e.In {
// 				// 	items[i] = v.(int)
// 				// }
// 				// fn(items)
// 				// t.Errorf("#%d\nin:  %v\nout: %v", e.Count, e.In, items)
// 			}
// 		})
// 	}
// }
