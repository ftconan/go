package err_test

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargeThanHundredError = errors.New("n should be not large than 100")

func GetFibonacci(n int) ([]int, error) {
	//if n < 2 || n  > 100 {
	//	return nil, errors.New("n should be in [2, 100]")
	//}
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n  > 100 {
		return nil, LargeThanHundredError
	}

	fibList := []int{1, 1}

	for i :=2; i < n; i++ {
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			t.Log("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
