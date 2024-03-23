package binary

import (
    "fmt"
)

func Run() {
    fmt.Println("Running Binary Search...")
	numbers := []int{2, 8, 5, 3, 9, 4, 1}

    fmt.Println(binary_search(numbers, 3))
}

func binary_search(a []int, target int) int {
    left := 0
    right := len(a) - 1

    for left <= right {
        mid := (right + left) / 2

        if a[mid] == target {
            return mid
        } else if a[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
