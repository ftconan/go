// Author: magician
// File:   filter.go
// Date:   2020/4/12
package pipe_filter

type Request interface {

}

type Response interface {

}

type Filter interface {
	Process(data Request) (Response, error)
}
