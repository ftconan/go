// Author: magician
// File:   memo_test.go
// Date:   2020/7/30
package memo

import (
	memo "golang/go_program/ch9/memo1"
	"golang/go_program/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
