package worker

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"spider/crawler_distributed/config"
	"spider/engine"
	"spider/model"
	"spider/webs/zongheng/parser"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items   []engine.Item
	Request []Request
}

//把engine里的Request转化成这边可以在网络上传输的Request
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

//把engine里的ParseResult转化成这边可以在网络上传输的ParseResult
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items:   r.Items,
		Request: nil,
	}
	for _, req := range r.Requests {
		result.Request = append(result.Request, SerializeRequest(req))
	}
	return result
}

//反序列化
func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, err
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Request {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializeing request:%v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

//反序列化成解析器 方法一，维护一个全局的map，提前把所有函数的名字和这个函数注册进去，方法二，这边使用的switch case
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCatalog:
		return engine.NewFuncParser(parser.ParseCatalog, config.ParseCatalog), nil
	case config.ParseNovelHomePage:
		return engine.NewFuncParser(parser.ParseNovelHomePage, config.ParseNovelHomePage), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseNovelContent:
		if n, ok := p.Args.(model.Novel); ok {
			return parser.NewNovelContentParser(n), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
