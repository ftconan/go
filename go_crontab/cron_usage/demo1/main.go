// Author: magician
// File:   main.go
// Date:   2020/5/5
package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)

	// 分（0-59） 时（0-23） 天（1-31） 月（1-12） 星期（0-6）

	// 每分钟执行一次
	//if expr, err = cronexpr.Parse("* * * * *"); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// 每隔5分钟执行一次
	if expr, err = cronexpr.Parse("*/6 * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 当前时间
	now = time.Now()
	// 下次调度时间
	nextTime = expr.Next(now)
	fmt.Println(now, nextTime)
}
