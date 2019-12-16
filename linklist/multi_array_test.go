package linklist

import(
	"testing"
	"fmt"
)


func TestAll(t *testing.T) {

	var L Pointer = -1

	for i := 0; i < 5; i++ {
		p := Aolloc()
		p.SetData(i)

		p.SetNext(L)
		p.SetPrev(-1)

		if L != -1 {
			L.SetPrev(p)
		}
		L = p
	}

	p := L
	for true {
		if p == -1 {
			break;
		}
		fmt.Println(p.Data())
		p = p.Next()
	}	
}
