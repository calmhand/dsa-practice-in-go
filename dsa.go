package main

import (
	"fmt"
	bubbleSort "remo/bubble"
    insertionSort "remo/insertion"
    selectionSort "remo/selection"
    quickSort "remo/quick"
    mergeSort "remo/merge"
    heapSort "remo/heap"
)

func main() {
	fmt.Println("DSA Practice")
	bubbleSort.Run()
    insertionSort.Run()
    selectionSort.Run()
    quickSort.Run()
    mergeSort.Run()
    heapSort.Run()
}
