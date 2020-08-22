package Sort

import (
	"testing"
)

func TestSelectSort(t *testing.T){

	inputArray := [][]int{
		{1, 23, 25, 29, 12, 8, 4, 100},
		{2, 12, 88, 33, 5, 9, 1, 22},
		{-1},
		{0},
		{11, 22 ,33 ,44 ,55, 66, 77, 88, 99},
	}
	for _, val := range inputArray{
		_ , err := SelectSort(val)
		if err != nil{
			t.Fatalf("[TestSelectSort] Fail to call ! err:%v\n", err)
		}
	}

	t.FailNow()

}
