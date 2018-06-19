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
)

func httpGet(url string,params string)string  {

	if len(url)<=0 {
		return ""
	}

	targetUrl :=url+"?"+params

	resp, err :=   http.Get(targetUrl)

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