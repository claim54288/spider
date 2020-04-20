package engine

import "github.com/PuerkitoBio/goquery"

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type Request struct {
	Url        string
	ParserFunc func(document *goquery.Document) ParseResult
}

//空解析器，不做事情，防止nil报错
func NilParser(*goquery.Document) ParseResult {
	return ParseResult{}
}
