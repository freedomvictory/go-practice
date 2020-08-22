package Sort

import (
	"fmt"
	"log"
)

func SelectSort(inputArray []int) ([]int, error){

	if inputArray == nil || len(inputArray) == 0{
		return nil, fmt.Errorf("[SelectSort] (error) The input array can't nil or empty")
	}

	var max, max_index int = inputArray[0], 0
	sortedArray := []int{}
	arrayLen := len(inputArray)

	for i := 0; i < arrayLen; i++{
		for j := i; j < arrayLen; j++{
			if inputArray[j] > max{
				max = inputArray[j]
				max_index = j
			}
		}
		sortedArray = append(sortedArray, max)
		max = 0
		inputArray[max_index] = inputArray[i]
	}
	log.Printf("[SelectSort] (log) outputArray:%v\n", sortedArray)
	return sortedArray, nil
}



