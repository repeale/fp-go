package fp

import (
	"fmt"
	"testing"
)

func TestConverge_Example(t *testing.T) {
	concatNumbers := func(n1, n2 int) string {
		return fmt.Sprintf("%d%d", n1, n2)
	}
	res := Converge2(concatNumbers, add1, double)(10)
	expected := "1120"

	if res != expected {
		t.Errorf("Should converge results of two functions. Expected: %s, Received:%s", expected, res)
	}
}
