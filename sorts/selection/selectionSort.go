package selection

import (
    "fmt"
)

// During each iteration, select the smallest item from the unsorted partition and move it to the sorted partition.
// The "current minimum" and "current item" are tracked.

// Worst Case: O(n^2)
// Best Case: O(n)

func Run() {
    fmt.Println("Running Selection Sort...")
    numbers := [7]int{2, 8, 5, 3, 9, 4, 7}

    fmt.Println(fmt.Sprintf("%s%d", "Before: ", numbers))

    for i := 0; i < len(numbers) - 1; i++ {
        cur_min := i

        for j := i + 1; j < len(numbers); j++ {
            if numbers[j] < numbers[cur_min] {
                cur_min = j
            }
        }

        if cur_min != i {
            numbers[i], numbers[cur_min] = numbers[cur_min], numbers[i]
        }
    }

    fmt.Println(fmt.Sprintf("%s%d\n", "After: ", numbers))
}
