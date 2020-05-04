// Author: magician
// File:   main.go
// Date:   2020/5/4
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	// 生成cmd
	cmd = exec.Command("/bin/bash", "-c", "sleep 1;ls -l;")

	// 执行了命令，捕获了子进程的输出（pipe）
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	// 打印子进程的输出
	fmt.Println(string(output))
}
