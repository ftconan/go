// Author: magician
// File:   to_int_filter.go
// Date:   2020/4/12
package pipe_filter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToIntFilter struct {

}

func NewToIntFilter() *ToIntFilter  {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error)  {
	// 数据类型转换
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}
