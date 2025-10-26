package main

// SortingAlgo defines how you want to sort your data.
// Each strategy implements this differently.
type SortingAlgo interface {
	sort(s *Sorter)
}