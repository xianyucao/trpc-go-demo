package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/proto/simplest"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	simplest.RegisterHelloWorldService(s, helloWorldImpl{})
	_ = s.Serve()
}

type helloWorldImpl struct{}

func (helloWorldImpl) Hello(ctx context.Context, req *simplest.HelloRequest) (*simplest.HelloResponse, error) {
	rsp := &simplest.HelloResponse{}
	res := req.Greeting
	var siteMap map[string]string
	siteMap = make(map[string]string)
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"
	name, ok := siteMap[res]
	var out string
	if ok {
		out = name
	} else {
		out = " nil "
	}
	out = out + " "
	rsp.Response = fmt.Sprintf("%s is %s", res, out)
	rsp.Timestamp = float64(time.Now().UnixMilli()) / 1000
	return rsp, nil
}
