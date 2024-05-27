package sorts

func MergeSort[T Datatype](Data []T) []T{
	if len(Data)<=1{
		return Data
	}
	med := len(Data)/2
	left := MergeSort(Data[:med])
	right := MergeSort(Data[med:])


	return merge(left,right)
}

func merge[T Datatype](left,right []T) []T{
	var i,j int
	var result []T
	for i < len(left) && j < len(right){
		if left[i]<right[j]{
			result = append(result, left[i])
			i++
		}else{
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}



