package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 访问URL地址，并拿到所有返回的数据

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bytes, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println("error: status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d wrong meg:%s", resp.StatusCode, string(bytes))
	}

	return ioutil.ReadAll(resp.Body)
	/*all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)
	printCityList(all)*/
}
