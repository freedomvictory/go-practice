package Sort



func Qsort(array []int) []int{

	if len(array) < 2{
		return array
	}else{
		RefVal := array[0]
		leftArray := []int{}
		rightArray := []int{}
		for i := 1; i < len(array); i ++{
			if array[i] > RefVal{
				rightArray = append(rightArray, array[i])
			}else{
				leftArray = append(leftArray, array[i])
			}
		}

		retArray := []int{}
		retArray = append(retArray, Qsort(leftArray)...)
		retArray = append(retArray, RefVal)
		retArray = append(retArray, Qsort(rightArray)...)

		return retArray

	}







}


