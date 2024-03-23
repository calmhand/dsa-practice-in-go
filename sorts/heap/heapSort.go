package heap

import (
	"fmt"
)

// Some code will be referenced from a repository by Dan Rusei.

// A heap is a ordered binary tree. A 'max heap' has a strict rule where the value of the
// parent nodes are of a larger value than the child nodes (vice-versa for a 'min heap').
// Steps: Create a max heap, remove the largest item, and place the item in a sorted partition.
// 

// Worst Case: O(n*logn)
// Best Case: O(n*logn)

func Run() {
	fmt.Println("Running Heap Sort...")
	numbers := []int{2, 8, 5, 3, 9, 4, 7} // If the length of an array is odd, the tree is "complete". 

	fmt.Println(fmt.Sprintf("%s%d", "Before: ", numbers))

	heapSort(numbers, len(numbers))

	fmt.Println(fmt.Sprintf("%s%d\n", "After: ", numbers))
}

func heapSort(a []int, length int) {
	for i := length/2 - 1; i >= 0; i-- { // i == index of parent node
		heapify(a, length, i)
	}

	for i := length - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, i, 0)
	}
}

func heapify(a []int, length, i int) { // O(n)
	parent := i           // index of a parent node
	leftChild := 2*i + 1  // index of the left child of parent
	rightChild := 2*i + 2 // index of the right child of parent

	if leftChild < length && a[leftChild] > a[parent] { // If the index of the child is NOT sorted, AND the left child is larger than its parent.
		parent = leftChild
	}

	if rightChild < length && a[rightChild] > a[parent] {
		parent = rightChild
	}

	if parent != i { // Check if a new value was to parent from previous if conditions.
		a[i], a[parent] = a[parent], a[i]
		heapify(a, length, parent) // heapify the sub-tree after the swap
	}
}
