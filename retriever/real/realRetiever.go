package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retrieverrr struct {
	UserAgent string
	TimeOut   time.Duration
}

// 指针接受者，通过这种接受指针的方式，实现 interface 中的函数
func (r *Retrieverrr) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(result)
}
