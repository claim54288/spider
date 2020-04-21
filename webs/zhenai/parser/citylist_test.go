package parser

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun") //测试用时候就不用当场拉了，直接把网页信息存本地
	const filename = "citylist_test_data.html"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	contents, _ := goquery.NewDocumentFromReader(file)
	//fmt.Println(contents.Html())
	result := ParseCityList(contents)

	//verify result 验证结果
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("结果数量不对，应该是%d，实际是%d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("item结果数量不对，应该是%d，实际是%d", resultSize, len(result.Items))
	}
}
