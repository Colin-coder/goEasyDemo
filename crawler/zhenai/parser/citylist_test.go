package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	/*
		contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	*/

	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)

	result := ParserCityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("res should have %d "+"req; but had %d", resultSize, len(result.Items))
	}

	fmt.Println(result)
}
