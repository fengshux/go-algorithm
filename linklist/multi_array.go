package linklist



var (
	next = make([]int, 1024)
	prev = make([]int, 1024)
	data = make([]interface{}, 1024)
	free Pointer
)

type Pointer int

func(p Pointer) next() int {
	return next[p]
}





func init() {
	free = 0;
	for i:=0; i < 1023; i++ {
		next[i] = i+1
	}
	next[1023] = -1
}
