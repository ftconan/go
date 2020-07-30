// Author: magician
// File:   memo_test.go
// Date:   2020/7/30
package memo_test

import (
	memo "golang/go_program/ch9/memo5"
	"golang/go_program/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
