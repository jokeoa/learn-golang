package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1)
	}
	return binarySearch(arr, target, mid+1, right)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	fmt.Println(factorial(5))
	
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println(binarySearch(arr, 7, 0, len(arr)-1))
	
	sortArr := []int{64, 34, 25, 12, 22, 11, 90}
	quickSort(sortArr, 0, len(sortArr)-1)
	fmt.Println(sortArr)
}
