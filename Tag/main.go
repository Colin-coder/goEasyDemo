package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
golang 中 tag 标签的使用方法
*/

// 这里 tag 部分编写有语法问题，label:前后不能加任何空格，uppercase:前可加一个空格，后不能加空格，也不能写成中间加空格的 label :
type Person struct {
	Name        string `label:"Name is :" uppercase:"true"`
	Age         int    `label:"Age is :"`
	Description string
}

func Func1() {

	person := Person{Name: "nk", Age: 10, Description: "ddd"}

	//fmt.Println(person)

	err := PrintTag(&person)
	if err != nil {
		fmt.Println(err)
	}
}

func PrintTag(ptr interface{}) error {
	t := reflect.TypeOf(ptr) // 取出变量类型
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		// 如果变量的类型不是指针，或者变量的
		return fmt.Errorf("param is invalid, param : %v", t)
	}

	// reflect.ValueOf(ptr) 取出变量 ptr 的类型，这里是指针类型
	// Elem() 取出 ptr 指针实际指向的值
	v := reflect.ValueOf(ptr).Elem()

	// NumField，取出struct中的元素数量，进行遍历
	for i := 0; i < v.NumField(); i++ {
		// 取出结构体的每个成员元素
		fieldInfo := v.Type().Field(i) // 取出v.Type()这个值中的类型，.Field(i)再取出这个类型中的第i个域（这里研究的是类型的域）
		tag := fieldInfo.Tag           // 取出类型中这个域的tag值

		label := tag.Get("label")
		if label == "" {
			// fieldInfo.Name 域的名称，这里是指变量名称
			label = fieldInfo.Name + ": "
		}

		value := fmt.Sprintf("%v", v.Field(i))       // 把v.Field(i)中第i个域转换为string类型
		if fieldInfo.Type.Kind() == reflect.String { // 判断这个field是否是string变量类型
			uppercase := tag.Get("uppercase")
			if uppercase == "true" {
				value = strings.ToUpper(value)
			} else {
				value = strings.ToLower(value)
			}
		}

		fmt.Println(label + value)
	}
	return nil
}

func main() {
	Func1()
}
