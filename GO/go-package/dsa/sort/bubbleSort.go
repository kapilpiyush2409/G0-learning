package sorts

func BubbleSort[T Datatype](data *[]T) {
	size := len(*data)
	for i, _ := range *data {
		for j := 0; j < size-1-i; j++ {
			if (*data)[j] > (*data)[j+1] {
				(*data)[j], (*data)[j+1] = (*data)[j+1], (*data)[j]
			}
		}
	}
}

// Time Complexity
// Best	O(n)
// Worst	O(n2)
// Average	O(n2)
// Space Complexity	O(1)
// Stability
