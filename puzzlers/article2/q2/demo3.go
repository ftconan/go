// Author: magician
// File:   demo3.go
// Date:   2021/1/7
package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("", flag.ExitOnError)

func init() {
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	//flag.CommandLine.Usage = func() {
	//	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s: \n", "question")
	//	flag.PrintDefaults()
	//}
	////flag.StringVar(&name, "name", "everyone", "The greeting object.")

	// cmdLine
	cmdLine.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s: \n", "question")
		cmdLine.PrintDefaults()
	}
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	//flag.Usage = func() {
	//	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s: \n", "question")
	//	flag.PrintDefaults()
	//}
	//flag.Parse()
	_ = cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello %s!\n", name)
}
