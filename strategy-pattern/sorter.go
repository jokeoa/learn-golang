package main

import "fmt"

// Sorter holds your data and the strategy you're using to sort it.
// You can swap strategies at runtime without changing your data.
type Sorter struct {
	data         []int
	sortingAlgo  SortingAlgo
}

// initSorter creates a new Sorter with your chosen strategy.
// Pass in the data you want to sort and how you want to sort it.
func initSorter(data []int, algo SortingAlgo) *Sorter {
	return &Sorter{
		data:        data,
		sortingAlgo: algo,
	}
}

// setSortingAlgo lets you change strategies on the fly.
// This is the whole point of the pattern - flexible behavior.
func (s *Sorter) setSortingAlgo(algo SortingAlgo) {
	s.sortingAlgo = algo
}

// sort delegates to whatever strategy you've chosen.
// The Sorter doesn't care how it's done, just that it gets done.
func (s *Sorter) sort() {
	s.sortingAlgo.sort(s)
}

// printData shows your array so you can see what's happening.
func (s *Sorter) printData() {
	fmt.Println(s.data)
}