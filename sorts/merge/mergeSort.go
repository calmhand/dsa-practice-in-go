package merge

import (
    "fmt"
)

// Merge sort is a divide and conquer algorithm.

// The initial array will be divided into two equal halves until a sub-array 
// of size 1 remains. Then, the two halves will be sorted recursively.
// Two functions are involved: mergeSort and merge.


// Worst Case: O(n*logn)
// Best Case: O(n*logn)

func Run() {
    fmt.Println("Running Merge Sort...")
    numbers := []int{2, 8, 5, 3, 9, 4, 7}

    fmt.Println(fmt.Sprintf("%s%d", "Before: ", numbers))

    numbers = mergeSort(numbers)  

    fmt.Println(fmt.Sprintf("%s%d\n", "After: ", numbers))

}

func mergeSort(a []int) []int {
    // Base case
    if len(a) < 2 {
        return a
    }
    
    // "a[:len(a) / 2]" is an example of Slices in go.
    // Ex: a[low : high]
    // This syntax is basically saying: Slice off the last "len(a) / 2" elements and return an
    // array with the rest of the elements.
    first := mergeSort(a[:len(a) / 2]) // Call mergeSort on the first half of the array.
    second := mergeSort(a[len(a) / 2:]) // Call mergeSort on the second half of the array.

    return merge(first, second) // Once we're at the lowest point, begin the merge.
}

func merge(a, b []int) []int {
    final := []int{}
    i := 0
    j := 0

    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }

    for ; i < len(a); i++ {
        final = append(final, a[i])
    }

    for ; j < len(b); j++ {
        final = append(final, b[j])
    }

    return final
}
