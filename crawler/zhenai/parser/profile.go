package parser

import (
	"goEasyDemo/crawler/engine"
	"goEasyDemo/crawler/model"
	"regexp"
	"strconv"
)

var idRex = regexp.MustCompile(`<div class="id" data-v-5b109fc3>ID：([\d]+)</div>`)

var heightReg = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)

func ParseProfile(contents []byte, name string) engine.ParserResult {

	profile := model.Profile{}
	profile.Name = name
	profile.Id = extractString(contents, idRex)

	profile.Height, _ = strconv.Atoi(extractString(contents, heightReg))

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result

}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
