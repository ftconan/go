package my_func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T)  {
	//a, b := returnMultiValues()
	//t.Log(a, b)
	a, _ := returnMultiValues()
	t.Log(a)

	ts := timeSpent(slowFun)
	t.Log(ts(10))
}
