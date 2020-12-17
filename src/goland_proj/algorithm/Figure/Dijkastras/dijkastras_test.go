package Dijkastras

import (
	"testing"
)

func TestDjks(t *testing.T){

  i  := make(map[string]NodesAndDistances, 32)


  a_nodes := map[string]int{
  	"fin":1,
  }
  b_nodes := map[string]int{
  	"a" : 3,
  	"fin": 5,
  }

  start_nodes := map[string]int{
  	"a":6,
  	"b":2,
  }

  i["start"] = NodesAndDistances{
  	Value: start_nodes,
  }
  i["a"] = NodesAndDistances{
		Value: a_nodes,
  }
  i["b"] = NodesAndDistances{
		Value: b_nodes,
  }
  i["fin"] = NodesAndDistances{}

  dist, out_path :=  FasterestToDest(i)
  t.Logf("[TestDjks] (log) dist :%v, out_path:%v\n", dist, out_path )
  t.FailNow()

}
