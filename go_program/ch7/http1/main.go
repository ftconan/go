// Author: magician
// File:   demo1.go
// Date:   2020/7/10
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string  {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	for item, price := range db {
		_, _ = fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// handler
//type Handler interface {
//	ServeHTTP(w http.ResponseWriter, r *http.Request)
//}
//
//func ListenAndServe(address string, h Handler) error
