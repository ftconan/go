// Author: magician
// File:   struct_def.go
// Date:   2020/4/14
package jsontest

type BasicInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo JobInfo `json:"job_info"`
}
