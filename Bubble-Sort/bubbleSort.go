package bubbleSort

import "fmt"

// Compare consecutive items. If they are "out of place" or a[n] > a[n + 1], swap them.
// The highest number will "bubble" towards the right with each iteration.
// A sorted partition will form at the end of the array.

// Worst case: O(n^2)
// Best case: O(n)

/* Notes - - -
    > = Sorts in Ascending Order
    < = Sorts in Decending Order
*/

func Run() {
	fmt.Println("Running bubble sort...")
	numbers := [7]int{2, 8, 5, 3, 9, 4, 1}

    fmt.Println(fmt.Sprintf("%s%d", "Before: ", numbers))
    
    for i := 1; i < len(numbers); i++ {
        for j := 0; j < len(numbers) - 1; j++ {
            if (numbers[j] > numbers[j + 1]) {
                // swap
                temp := numbers[j]
                numbers[j] = numbers[j + 1]
                numbers[j + 1] = temp
            }
        }
    }

    fmt.Println(fmt.Sprintf("%s%d", "After: ", numbers))


	/* Pseudocode - - -
	   for (i from 1 to N)
	     for (j from 0 to N - 1)
	       if (a[j] > a[j + 1])
	         swap (a[j], a[j + 1])
	       else continue (this is implied)
	*/
}
