package main

import "fmt"

// BubbleSort is the simple, slow strategy.
// It compares neighbors and swaps them until everything's in order.
type BubbleSort struct {
}

// sort pretends to do bubble sort.
// In real life, this would repeatedly swap adjacent elements.
func (b *BubbleSort) sort(s *Sorter) {
	fmt.Println("Sorting with bubble sort - comparing neighbors, swapping as needed")
	fmt.Printf("Would process %d elements with O(n²) comparisons\n", len(s.data))
}