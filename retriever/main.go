package main

import (
	"fmt"
	"goEasyDemo/retriever/mock"
	"goEasyDemo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	res := s.Post(url, map[string]string{
		"contents": "another faked",
	})
	fmt.Printf("result: %v\n", res)

	return s.Get(url)
}

func inspect(r Retriever) {

	fmt.Println("inspecting:")
	fmt.Printf("%T, %v\n", r, r)
	switch v := r.(type) {
	case mock.Retrieveree:
		fmt.Println("contents:", v.Contents)
	case *real.Retrieverrr:
		fmt.Println("userAgent:", v.UserAgent)
	}
}

// 一个普通的b interface{} 包含两个成员变量，
// 1. 实现者的类型，通过 a.(type) 来获取
// 2. 实现者的指针，它指向实现者
func main() {
	var r Retriever

	// 值接受者
	rr := mock.Retrieveree{"this is a ac"}
	inspect(rr)

	// 指针接受者
	r = &real.Retrieverrr{
		UserAgent: "moasdf",
		TimeOut:   time.Minute,
	}

	// 通过type switch 方法查看 interface 存储的类型  这个变量的值
	inspect(r)

	// Type assertion  方法查看 interface 存储的类型  这个变量的值
	realRetriever, ok := r.(*real.Retrieverrr)
	fmt.Println(ok)
	fmt.Printf("format: %T %v\n", realRetriever, realRetriever)
	fmt.Println(realRetriever.TimeOut)
	// realRetriever  它的类型是   *real.Retrieverrr
	// 值是：&{moasdf 1m0s}

	//io.ReadWriteCloser()
	/*
		mocRetriever := r.(mock.Retrieveree)
		fmt.Printf("%T\n", mocRetriever)
		fmt.Println(mocRetriever.Contents)
	*/
	//fmt.Println(download(r))

	fmt.Println("try a session")
	fmt.Println(session(rr))
}
