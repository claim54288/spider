package engine

import "github.com/PuerkitoBio/goquery"

type ParserFunc func(*goquery.Document) ParseResult

type Parser interface {
	Parse(*goquery.Document) ParseResult
	Serialize() (name string, args interface{})
}

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
	Url    string
	Parser Parser
}

type NilParser struct {
}

func (NilParser) Parse(*goquery.Document) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

//空解析器，不做事情，防止nil报错
//func NilParser(_ *goquery.Document) ParseResult {
//	return ParseResult{}
//}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents *goquery.Document) ParseResult {
	return f.parser(contents)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
