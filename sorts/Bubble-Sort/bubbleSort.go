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
	fmt.Println("Running Bubble Sort...")
	numbers := [7]int{2, 8, 5, 3, 9, 4, 1}

    fmt.Println(fmt.Sprintf("%s%d", "Before: ", numbers))
    
    for i := 0; i < len(numbers); i++ { // As 'i' reaches n - 1, the array will be sorted.
        for j := 0; j < len(numbers) - i - 1; j++ { // After 'i' iterations, the last 'i' elements are already sorted.
            if (numbers[j] > numbers[j + 1]) {
                // swap
                numbers[j], numbers[j + 1] = numbers[j + 1], numbers[j]
            }
        }
    }

    fmt.Println(fmt.Sprintf("%s%d\n", "After: ", numbers))


	/* Pseudocode - - -
	   for (i from 0 to N)
	     for (j from 0 to N - i - 1)
	       if (a[j] > a[j + 1])
	         swap (a[j], a[j + 1])
	       else continue (this is implied)
	*/
}
