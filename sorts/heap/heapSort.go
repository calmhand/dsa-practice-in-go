package heap

import (
    "fmt"
)

func Run() {
    fmt.Println("Running Heap Sort...")
    numbers := [7]int{2, 8, 5, 3, 9, 4, 7}

    fmt.Println(fmt.Sprintf("%s%d", "Before: ", numbers))

    

    fmt.Println(fmt.Sprintf("%s%d", "After: ", numbers))
}
