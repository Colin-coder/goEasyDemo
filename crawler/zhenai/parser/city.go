package parser

import (
	"goEasyDemo/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

// 城市列表解析器，输入为从URL请求到的所有数据
func ParserCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)

	// 这里返回 [][][]byte 第一个[]代表是一个列表，第二个[]代表列表中每个元素都是一个[]byte（可以近似看成string）类型的数组，且每个数组都是只有两个元素
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		// Items 人的名称 m[2]
		result.Items = append(result.Items, "User "+string(m[2]))

		// m[1] 每个人的URL地址
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc: func(bytes []byte) engine.ParserResult {
			//	return ParseProfile(contents, string(m[2]))
			//}, // 每个城市URL对应的解析器设置为nil
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
