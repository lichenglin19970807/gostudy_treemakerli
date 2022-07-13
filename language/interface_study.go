// Package language show my study of golang
package language

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	name string
}

type TrapError struct {
	code int
	msg  string
}

func (te *TrapError) Error() string {
	return te.msg
}

func returnNil() *TrapError {
	return nil
}

// 陷阱：通过接口修改数据
// 数据指针持有的是目标对象的只读复制品
func TrapOne() {
	u := User{1, "A"}
	var i interface{} = u
	u.id = 2
	u.name = "B"
	fmt.Printf("%v\n", u)        // {2 B}
	fmt.Printf("%v\n", i.(User)) // {1 A}

	var p interface{} = &u
	p.(*User).name = "C"
	fmt.Printf("%v\n", u)
}

func TrapTwo() {
	var err error = returnNil()
	if err == nil {
		fmt.Printf("err is nil")
		return
	}
	fmt.Println("err is not nil")
	fmt.Println(reflect.ValueOf(err).IsNil())
	fmt.Println(reflect.TypeOf(err))
}
