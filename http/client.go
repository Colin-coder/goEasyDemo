package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resq, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)

	resq.Header.Add("User-Agent",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36")

	nkclient := http.Client{

		// 重定向操作
		CheckRedirect: func(req *http.Request, via []*http.Request) error {

			fmt.Println("Redirect", req)
			return nil
		},
	}

	//resp, err := http.DefaultClient.Do(resq)
	resp, err := nkclient.Do(resq)
	//resp, err := http.Get()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, e := httputil.DumpResponse(resp, true)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%s \n", bytes)
}
