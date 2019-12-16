package linklist


var (
	next = make([]Pointer, 1024)
	prev = make([]Pointer, 1024)
	data = make([]interface{}, 1024)
	free Pointer
)

type Pointer int

func(p Pointer) Next() Pointer {
	return next[p]
}

func(p Pointer) Prev() Pointer{
	return prev[p]
}

func(p Pointer) Data() interface{} {
	return data[p]
}


func(p Pointer) SetNext(ne Pointer) {
	next[p] = ne
}

func(p Pointer) SetPrev(pr Pointer) {
	prev[p] = p
}

func(p Pointer) SetData(d interface{}) {
	data[p] = d
}


func Aolloc() Pointer{
	if free == -1 {
		panic("stack overflow")
	}
	
	p := free
	free = free.Next()
	return p
}

func Free(p Pointer) {
	next[p] = free
	free = p
}

func init() {
	free = 0;
	for i := 0; i < 1023; i++ {
		next[i] = Pointer(i+1)
	}
	next[1023] = -1
}
