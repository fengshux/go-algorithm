package sort

import(
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func TestCounterSort(t *testing.T) {
	var (
		arr []int
		max int = -1
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		item := rand.Intn(100) 
		arr = append(arr,item)
		if item > max {
			max = item
		}
	}
	
	sorted := CounterSort(arr, max)
	fmt.Println(sorted)

	for i := 1; i < len(sorted); i++ {
		if sorted[i-1] > sorted[i] {
			t.Error("Not sorted.")
		}
	}
}
