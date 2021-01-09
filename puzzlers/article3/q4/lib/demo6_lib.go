// Author: magician
// File:   demo6_lib.go
// Date:   2021/1/9
package lib

import (
	in "golang/puzzlers/article3/q4/lib/internal"
	"os"
)

func Hello(name string)  {
	in.Hello(os.Stdout, name)
}
