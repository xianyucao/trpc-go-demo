package main

import (
	"context"
	"fmt"
	"time"

	// proto package 的路径请读者自行调整

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
	rsp.Response = fmt.Sprintf("%s to you, too", req.Greeting)
	rsp.TimestampMsec = time.Now().UnixMilli()
	return rsp, nil
}
