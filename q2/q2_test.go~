package main
import "testing"

type rectestpair struct{
	recs Rectangle
	result float64
}
type cirtestpair struct{
	cirs Circle
	result float64
}

var rectests = []rectestpair{{Rectangle{0,0,2,2},8},{Rectangle{0,0,4,4},16}}
var cirtests = []cirtestpair{{Circle{0,0,4},25.132741228718345}}

func TestPerimeter(t *testing.T){
//Testing Rectangle Perimeter
	for _,cases:= range rectests{
		v1:=cases.recs.Perimeter()
		if v1!=cases.result{
			t.Error("Expected",cases.result, "got",v1)
		}
	}
//Testing Circle Perimeter
	for _,cases:= range cirtests{
		v1:=cases.cirs.Perimeter()
		if v1!=cases.result{
			t.Error("Expected",cases.result, "got",v1)
		}
	}
}
