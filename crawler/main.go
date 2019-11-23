package main

import (
	"goEasyDemo/crawler/engine"
	"goEasyDemo/crawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
