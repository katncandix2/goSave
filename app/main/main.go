/*
@Time : 2018/6/19 16:39 
@Author : lenovo
@File : main
@Software: GoLand
*/
package main

import (
	"goSave/app/Utils"
	"goSave/app/Constants"
	"encoding/json"
	"fmt"
)


func main()  {

	str :="stuIdArr[]=8336677"

	res := Utils.HttpPost(str,Constants.USER_ORDER_BY_USERID_ARR)

	var data interface{}

	json.Unmarshal([]byte(res),&data)

	fmt.Println(data)

	}