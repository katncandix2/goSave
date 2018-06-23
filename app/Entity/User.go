/*
@Time : 2018/6/13 21:36 
@Author : lenovo
@File : User
@Software: GoLand
*/
package Entity

type User struct {
	name string
	age  int

	value *superMan
}

type superMan interface {
	fly()
	jump()
	sayHello()
}


