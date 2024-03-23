package main

import (
	"fmt"
	bubbleSort "remo/bubble"
    insertionSort "remo/insertion"
    selectionSort "remo/selection"
    quickSort "remo/quick"
    mergeSort "remo/merge"
    heapSort "remo/heap" 

    binarySearch "remo/binary"
)

func main() {
    fmt.Println("DSA: Sorting Practice")
	bubbleSort.Run()
    insertionSort.Run()
    selectionSort.Run()
    quickSort.Run()
    mergeSort.Run()
    heapSort.Run()

    fmt.Println("DSA: Search Practice")
    binarySearch.Run()
}
