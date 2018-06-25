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

	userName := req.FormValue("name")

	fmt.Fprintln(res, "hello world"+userName)
}

func main() {
	http.HandleFunc("/hello", IndexHandler)
	//http.ListenAndServe("10.99.1.152:8080", nil)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
