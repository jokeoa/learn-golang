package main

import "fmt"

func main() {
	// Your data that needs sorting
	data := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Starting data:")
	fmt.Println(data)
	fmt.Println()

	// Try bubble sort first
	bubbleSort := &BubbleSort{}
	sorter := initSorter(data, bubbleSort)

	fmt.Println("--- Using Bubble Sort Strategy ---")
	sorter.sort()
	fmt.Println()

	// Switch to quick sort - same data, different strategy
	quickSort := &QuickSort{}
	sorter.setSortingAlgo(quickSort)

	fmt.Println("--- Switched to Quick Sort Strategy ---")
	sorter.sort()
	fmt.Println()

	// Switch to merge sort - again, same data, different approach
	mergeSort := &MergeSort{}
	sorter.setSortingAlgo(mergeSort)

	fmt.Println("--- Switched to Merge Sort Strategy ---")
	sorter.sort()
	fmt.Println()

	fmt.Println("Final data (unchanged in this demo):")
	sorter.printData()
}