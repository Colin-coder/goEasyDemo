package parser

import (
	"goEasyDemo/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 城市列表解析器，输入为从URL请求到的所有数据
func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)

	// 这里返回 [][][]byte 第一个[]代表是一个列表，第二个[]代表列表中每个元素都是一个[]byte（可以近似看成string）类型的数组，且每个数组都是只有两个元素
	matches := re.FindAllSubmatch(contents, -1)

	limit := 10
	result := engine.ParserResult{}
	for _, m := range matches {
		// Items 城市名称 m[2]
		result.Items = append(result.Items, "City "+string(m[2]))

		// m[1] 每个城市的URL地址
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity, // 每个城市URL对应的解析器设置为nil
		})
		limit--
		if limit < 0 {
			break
		}
		//fmt.Printf("City: %s, URL: %s\n ", m[2], m[1])
		//fmt.Println()
		//fmt.Printf("%s\n", m)
	}
	//fmt.Printf("matches num:%d\n", len(matches))
	return result
}
