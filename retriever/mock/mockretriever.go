package mock

import "fmt"

type Retrieveree struct {
	Contents string
}

func (r *Retrieveree) String() string {
	return fmt.Sprintf("Retrieveree :{contents=%s}", r.Contents)
}

// 值传递，这里不会改变 r.Contents
func (r Retrieveree) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

// 值接受者，通过值 (r Retrieveree) 接受，来实现interface 接口中的函数
func (r Retrieveree) Get(url string) string {
	return r.Contents
}
