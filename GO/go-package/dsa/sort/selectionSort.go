package sorts



func SelectionSort[T Datatype](data *[]T){
	for i, _ := range *data {
		min_step := i
		for j := i + 1; j < len(*data); j++ {
			if (*data)[j] < (*data)[min_step] {
				min_step = j
			}
		}
		(*data)[i], (*data)[min_step] = (*data)[min_step], (*data)[i]
	}
	
}


// Time Complexity	 
// Best	O(n2)
// Worst	O(n2)
// Average	O(n2)
// Space Complexity	O(1)
// Stability	No
