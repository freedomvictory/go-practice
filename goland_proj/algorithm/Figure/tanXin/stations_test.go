package tanXin

import (
	"fmt"
	"reflect"
	"testing"
)

//import "github.com/viant/toolbox"

import (
	"github.com/adam-hanna/arrayOperations"
)
func TestSetCover(t *testing.T){

	listOfAdvStations := map[string][]string{

		"kone":{
			"ID", "NV", "UT",
		},
		"ktwo":{
			"WA", "ID", "MT",
		},
		"kthree":{
			"OR", "NV", "CA",
		},
		"kfour":{
			"NV", "UT",
		},
		"kfive":{
			"CA", "AZ",
		},
	}

	stats_needed := []string{"MT", "WA", "OR", "ID", "NV", "UT"}

	out, err := SetCover(listOfAdvStations, stats_needed)
	if err != nil{
		t.Logf("TestSetCover error: %v\n", err)
		t.Fail()
	}
	fmt.Printf("[TestSetCover] out: %v\n", out)
	//t.FailNow()
}

func TestSetInterSection(t *testing.T){

	var a = []int{1, 1, 2, 3}
	var b = []int{1, 3, 4}

	z, ok := arrayOperations.Intersect(a, b)
	if !ok {
		fmt.Println("Cannot find intersect")
	}

	slice, ok := z.Interface().([]int)
	if !ok {
		fmt.Println("Cannot convert to slice")
	}
	fmt.Println(slice, reflect.TypeOf(slice))
}

func  TestSetDifference(t *testing.T)  {
	var a = []int{1, 2, 3, 4}
	var b = []int{1, 3}

	z, ok := arrayOperations.Difference(a, b)
	if !ok {
		fmt.Println("Cannot find difference")
	}

	slice, ok := z.Interface().([]int)
	if !ok {
		fmt.Println("Cannot convert to slice")
	}
	fmt.Printf("[TestSetDifference] result : %v", slice)



}