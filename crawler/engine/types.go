package engine

type Request struct {
	Url        string                    // 地址
	ParserFunc func([]byte) ParserResult // 解析器函数
}

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
