/*
@Time : 2018/6/19 17:20 
@Author : lenovo
@File : index
@Software: GoLand
*/
package main

import (
	"net/http"
	"fmt"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("----------------------------")
	fmt.Println(req)
	fmt.Println(req.FormValue("name"))
	fmt.Fprintln(res, "hello world")
}

func main() {
	http.HandleFunc("/hello", IndexHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
