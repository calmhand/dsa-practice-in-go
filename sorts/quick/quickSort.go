package quick

import (
    "fmt"
)

// Quicksort is a divide-and-conquer algorithm. It will make use of recursion.

// Worst Case: O(n^2)
// Best Case: O(n*logn)

func Run() {
    fmt.Println("Running Quick Sort...")
    numbers := []int{2, 8, 5, 3, 9, 4, 7}

    fmt.Println(fmt.Sprintf("%s%d", "Before: ", numbers))

    quicksort(numbers, 0, len(numbers) - 1)

    fmt.Println(fmt.Sprintf("%s%d\n", "After: ", numbers))
}

/*
    a - unsorted array
    start - first index of partitioned array
    end - last index of partitioned array
*/
func quicksort(a []int, start, end int) {
    if start < end {
        a, pivot := partition(a, start, end)
        quicksort(a, start, pivot - 1)
        quicksort(a, pivot + 1, end)
    }
}

func partition(a []int, start, end int) ([]int, int) {
    pivot := a[end] // Pick a pivot
    i := start // Set 'i' pointer to first index

    for j := start; j < end; j++ { // Traverse through the partitioned array
        if a[j] < pivot { // If the value at a[j] is less than the pivot, swap a[j] and a[i]
            a[j], a[i] = a[i], a[j]
            i++ // Increment 'i'
        }
    }

    a[i], a[end] = a[end], a[i] // swap a[hi] (the pivot), with a[i] (the sorted location of the pivot).

    return a, i
}
