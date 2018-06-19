/*
@Time : 2018/6/19 16:39 
@Author : lenovo
@File : main
@Software: GoLand
*/
package main

import (
	"goSave/app/Utils"
	"fmt"
)

func main()  {


	song := make(map[string]interface{})
	song["name"] = "李白"
	song["timelength"] = 128
	song["author"] = "李荣浩"

	res := Utils.HttpPost(song,"http://localhost:8080/hello")
	fmt.Println(res)
	}