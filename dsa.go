package main

import (
	"fmt"
	bubbleSort "remo/bubble"
    insertionSort "remo/insertion"
    selectionSort "remo/selection"
)

func main() {
	fmt.Println("DSA Practice")
	bubbleSort.Run()
    insertionSort.Run()
    selectionSort.Run()
}
