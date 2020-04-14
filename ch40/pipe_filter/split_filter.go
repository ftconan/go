// Author: magician
// File:   split_filter.go
// Date:   2020/4/12
package pipe_filter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter  {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error)  {
	// 数据类型转换
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
