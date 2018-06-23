/*
@Time : 2018/6/22 16:54 
@Author : lenovo
@File : FormateRes
@Software: GoLand
*/
package main

import (
	"goSave/app/Constants"
	"goSave/app/Utils"
	"github.com/tidwall/gjson"
	"fmt"
	"encoding/json"
)

type OrderInfo struct {
	Id 			interface{} `json:"id"`
	Order_num 	interface{}	`json:"Order_num"`
	User_id     interface{} `json:"user_id"`
	Real_price  interface{} `json:"real_price"`
	Status      interface{} `json:"status"`
	Order_type  interface{} `json:"order_type"`
	Create_time interface{} `json:"create_time"`
}

func main()  {

	str :="stuIdArr[]=8336677"

	res := Utils.HttpPost(str,Constants.USER_ORDER_BY_USERID_ARR)

	data := gjson.Get(res,"data.orderInfo.*")

	list := gjson.Get(res,"data.orderDetail.*")

	fmt.Println("####################")
	fmt.Println(list)


	orderInfo := data.String()

	fmt.Println(orderInfo)

	o := &OrderInfo{}
	o.Id =   gjson.Get(orderInfo,"id").String()
	o.User_id = gjson.Get(orderInfo,"stu_id").String()
	o.Status = gjson.Get(orderInfo,"status").String()
	o.Create_time = gjson.Get(orderInfo,"create_time").String()
	o.Order_type = gjson.Get(orderInfo,"order_type").String()
	o.Real_price = gjson.Get(orderInfo,"real_price").String()
	o.Order_num = gjson.Get(orderInfo,"order_num").String()

	fmt.Println(o)


	fmt.Println("-----------------")
	fmt.Println(*o)
	a,_ := json.Marshal(*o)

	fmt.Println(string(a))


	}