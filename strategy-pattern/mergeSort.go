package main

import "fmt"

// MergeSort is the stable, predictable strategy.
// It splits everything in half, sorts each half, then merges them.
type MergeSort struct {
}

// sort pretends to do merge sort.
// In real life, this would recursively split and merge arrays.
func (m *MergeSort) sort(s *Sorter) {
	fmt.Println("Sorting with merge sort - splitting in half, merging sorted pieces")
	fmt.Printf("Would process %d elements with O(n log n) guaranteed time\n", len(s.data))
}