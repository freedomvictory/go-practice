package breadthFirstSearch

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T){

	f := NewFigure()
	f.Items["you"] = []string{"alice", "bob", "claire"}
	f.Items["bob"] = []string{"anuj", "peggy"}
	f.Items["alice"] = []string{"peggy"}
	f.Items["claire"] = []string{"thom", "jonny"}
	f.Items["anuj"] = []string{}
	f.Items["peggy"] = []string{}
	f.Items["thom"] = []string{}
	f.Items["jonny"] = []string{}


	out := Search(f)
	t.Logf("[TestBreadthFirstSearch] (log) out:%v\n", out)
	t.FailNow()
}


