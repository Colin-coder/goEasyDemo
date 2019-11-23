package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) HelloFunc() {
	fmt.Println("hello")
}

func main() {
	//u := User{1, "OK", 15}
	//Info(u)

	m := Manager{User{1, "Ok", 16}, "oop"}

	// 取出变量的类型来研究
	t := reflect.TypeOf(m)

	// Anonymous:true 结果中这个字段反应是否是匿名字段
	fmt.Printf("%#v\n", t.Field(0)) // 取出User成员
	fmt.Printf("%#v\n", t.Field(1)) // 取出title成员

	// 这里的[]int{0, 0} 表示获取第0个field中的第0个元素，其他情况依次类推
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0})) // 取出User中的Id成员
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1})) // 取出User中的Name成员

	// 取出变量的值来研究
	val := reflect.ValueOf(&m)
	fmt.Printf("value of : %v \n", val)

	//val.Elem() 返回 ptr 实际指向的值，取出后可进行修改
	//Elem() 要求，value 必须是 ptr 或者 interface 类型，否则会panic错误
	fmt.Printf("value elem of : %v \n", val.Elem())

	// 通过reflect修改值
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(888)

	fmt.Println(x)
}

/*
Type: User
Fields:
    Id: int = 1
  Name: string = OK
   Age: int = 15
HelloFunc: func(main.User)
*/
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	// 判断输入的参数是否是结构体类型，而不是指针等其他类型
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("Param is not struct")
		return
	}

	// 打印User的每一个成员字段名称、类型、值
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	// 打印方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
