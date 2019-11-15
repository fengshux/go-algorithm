package sort

import (
	"testing"
	"math/rand"
	"fmt"
)

var queue *MaxPriorityQueue

func testInsert(t *testing.T) {
	length := rand.Intn(20) + 1
	for i:=0; i < length; i++ {
		v := rand.Intn(100)
		queue.Insert(v)
	}

	if queue.Size != length {
		t.Error("insert, size not equal")
	}
	
	for i := 1; i < (len(queue.Heap)-1/2); i++ {
		l := leftChaild(i)
		r := rightChild(i)

		if l < len(queue.Heap) && queue.Heap[i] < queue.Heap[l]{
			t.Error("insert, parent is smaller then leftchaild")
		}

		if r < len(queue.Heap) && queue.Heap[i] < queue.Heap[r]{
			t.Error("insert, parent is smaller then rightchaild")
		}
	}
	fmt.Println("testinsert", queue)	
}

func testMaxMum(t *testing.T) {
	max := queue.MaxMum()

	if queue.Heap[0] != max {
		t.Error("first is not Max")
	}
	
	for i := 1; i < len(queue.Heap); i++ {
		if queue.Heap[i] > max {
			t.Error("not return max")
		}
	}
	fmt.Println("testMaxMum", queue)
}

func testExtractMaxMum(t *testing.T) {
	oldSize := queue.Size
	max := queue.ExtractMaxMum()

	if oldSize == queue.Size {
		t.Error("MaxMum is not removed")
	}
	
	for i := 0; i < len(queue.Heap); i++ {
		if queue.Heap[i] > max {
			t.Error("not return max")
		}
	}
	fmt.Println("testExtractMaxMum", queue)
}

func testIncrease(t *testing.T) {
	i := rand.Intn(queue.Size)
	queue.Increase(i, rand.Intn(100))

	for i := 1; i < (len(queue.Heap)-1/2); i++ {
		l := leftChaild(i)
		r := rightChild(i)

		if l < len(queue.Heap) && queue.Heap[i] < queue.Heap[l]{
			t.Error("increase, parent is smaller then leftchaild")
		}

		if r < len(queue.Heap) && queue.Heap[i] < queue.Heap[r]{
			t.Error("increase, parent is smaller then rightchaild")
		}
	}
	fmt.Println("testIncrease", queue)
}

func TestAll(t *testing.T) {
	queue = NewMaxPriorityQueue()
	
	t.Run("TestInsert", testInsert)
	t.Run("TestMaxmum", testMaxMum)
	t.Run("TestExtractMaxmum", testExtractMaxMum)
	t.Run("TestIncrease", testIncrease)
}
