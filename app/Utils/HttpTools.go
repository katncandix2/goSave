/*
@Time : 2018/6/19 15:33 
@Author : lenovo
@File : HttpTools
@Software: GoLand
*/
package Utils

import (
	"net/http"
	"io/ioutil"
	"log"
	"goSave/app/Constants"
	"fmt"
	"unsafe"
	"encoding/json"
	"bytes"
)

func HttpGet(url string,params string)string  {

	if len(url)<=0 {
		return ""
	}

	targetUrl :=Constants.HTTP + url+"?"+params

	resp, err := http.Get(targetUrl)

	if err != nil {
		log.Print(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
	}

	return string(body)
}

func HttpPost(params map[string]interface{},url string)string  {

	bytesData, err := json.Marshal(params)

	if err != nil {
		fmt.Println(err.Error() )
		return ""
	}

	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str
}