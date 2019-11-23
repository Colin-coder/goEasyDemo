package engine

import (
	"goEasyDemo/crawler/fetcher"
	"log"
)

/*
引擎，负责不断循环，处理种子任务
*/

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, e := worker(r)
		if e != nil {
			continue
		}

		// 这里后面的 ... 代表的是将所有元素都append进去
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParserResult, error) {

	log.Printf("Fetching %s", r.Url)
	// fetcher 负责从URL地址中获取返回结果，文本数据
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v",
			r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
