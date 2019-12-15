package bitmap



import (
	"math"
)

var store []uint64


func powUint64(x, y uint16) uint64 {	
	return uint64(math.Pow(float64(x), float64(y)))
}

func Set(byt uint16, val uint8) {
	
	

}


func SetOne(byt uint16) {

	index := byt/64
	
	bt := byt%64
	
	count := powUint64(2, bt)

	store[index] = store[index]| count
}

func SetZero(byt uint16) {

	index := byt/64
	bt := byt%64

	count := ^powUint64(2, bt)
	store[index] = store[index]&count	

}


func Get(byt uint16) int64 {
	
	index := byt/64
	bt := byt%64
	count := powUint64(2, bt)
	val := store[index]&count

	if val == 0 {
		return 0
	}
	return 1
	
}

func init(){
	store = make([]uint64, 65536)
}
