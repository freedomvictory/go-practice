package recursion

import "testing"

func TestFib(t *testing.T){

	ns := []int{1, 2, 3, 4, 5 }
	for _, val := range ns{
		out := FibonacciSequence(val)
		t.Logf("[TestFib] out:%v\n", out)
	}

	t.FailNow()

}