// Author: magician
// File:   sum_filter.go
// Date:   2020/4/12
package pipe_filter

import (
	"errors"
)

var SumFilterWrongFormatError = errors.New("input data should be []int")

type SumFilter struct {

}

func NewSumFilter() *SumFilter  {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error)  {
	// 数据类型转换
	elems, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
