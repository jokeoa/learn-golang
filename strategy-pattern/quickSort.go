package main

import "fmt"

// QuickSort is the fast, divide-and-conquer strategy.
// It picks a pivot and partitions around it.
type QuickSort struct {
}

// sort pretends to do quick sort.
// In real life, this would pick pivots and recursively partition.
func (q *QuickSort) sort(s *Sorter) {
	fmt.Println("Sorting with quick sort - picking pivots, partitioning recursively")
	fmt.Printf("Would process %d elements with O(n log n) average time\n", len(s.data))
}