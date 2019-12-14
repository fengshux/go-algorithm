package bitmap

import(
	"testing"
	"fmt"
)



func testSetOne(t *testing.T) {
	
}


func testSetZero(t *testing.T) {
	
	
}


func testGet(t *testing.T) {

	
}


func TestAll(t *testing.T) {

	SetOne(64)

	for i := 0 ; i< 3; i++ {
		fmt.Println(store[i])
	}

	val := Get(64)

	fmt.Println(val)

	val = Get(16)
	fmt.Println(val)

	SetOne(15)

	for i := 0 ; i< 3; i++ {
		fmt.Println(store[i])
	}

	val = Get(15)
	fmt.Println(val)
}
