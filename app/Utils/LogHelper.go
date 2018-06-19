/*
@Time : 2018/6/19 15:39 
@Author : lenovo
@File : LogHelper
@Software: GoLand
*/
package Utils

import (
	"goSave/app/Constants"
	"os"
	"log"
)

//记录日志
func SaveLog(logName string,logLevel string,logContents string)	  {

	if len(logName)<=0 {
		return
	}

	fileName :=  Constants.LOG_PREFIX + logName + Constants.LOG_SUFFIX

	logFile,err  := os.Create(fileName)

	if err!=nil {
		//todo 待优化
		println(err.Error())
	}

	defer logFile.Close()

	logM := log.New(logFile,logLevel,log.Llongfile)

	logM.Println(logContents)

}

