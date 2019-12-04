package dp

import (
	"testing"
	"fmt"
)

func TestChangeWord (t *testing.T) {
	md := ChangeWord("abdc", "bd")
	if md !=2 {
		t.Error("Do not right.")
	}
	fmt.Println(md)
}
