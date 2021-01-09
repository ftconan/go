// Author: magician
// File:   internal.go
// Date:   2021/1/9
package internal

import (
	"fmt"
	"io"
)

func Hello(w io.Writer, name string)  {
	_, _ = fmt.Fprintf(w, "Hello, %s\n", name)
}
