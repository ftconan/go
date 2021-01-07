// Author: magician
// File:   demo1.go
// Date:   2020/5/4
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	cmd = exec.Command("/bin/bash", "-c", "echo 1;echo 2;")
	err = cmd.Run()
	fmt.Println(err)
}
