package insertion

import (
    "fmt"
)


// Starting from left to right, examine each item and compare it to items on its left.
// Insert the item in the correct position in the array.
// The array will form sorted and unsorted paritions.

// Worst case: O(n^2)
// Best case: O(n)

func Run() {
    fmt.Println("Running Insertion Sort...")
    numbers := [7]int{2, 8, 5, 3, 9, 4, 7}

    fmt.Println(fmt.Sprintf("%s%d", "Before: ", numbers))
    
    for i := 1; i < len(numbers); i++ {
        j := i
        for j > 0 && numbers[j - 1] > numbers[j] {
            numbers[j], numbers[j - 1] = numbers[j - 1], numbers[j]
            j--
        }
    }

    fmt.Println(fmt.Sprintf("%s%d\n", "After: ", numbers))
}
