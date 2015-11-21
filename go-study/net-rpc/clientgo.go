package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call
// 参数列表: serviceMethod 服务名称, args 发送的参数, reply 回复, done 一个Call类型的channel返回值, call 返回一个Call
// 这个函数用来异步调用RPC服务

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return fmt.Errorf("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	// 设置服务端
	arith := new(Arith)
	// 注册服务
	rpc.Register(arith)
	//将服务注册到HTTP协议
	rpc.HandleHTTP()
	// 创建连接监听
	l, err := net.Listen("tcp", ":1234")

	if err != nil {
		fmt.Println("listen error: ", err.Error())
		return
	}

	// 设置监听连接
	go rpc.Accept(l)
	// 设置自定义的servercodec

	// 暂停两秒让服务有足够的时间开启
	time.Sleep(time.Second * 2)

	//使用TCP 方式进行网络请求
	client, _ := rpc.Dial("tcp", "127.0.0.1:1234")
	defer client.Close()

	args := &Args{7, 8}
	var reply int
	call := client.Go("Arith.Multiply", args, &reply, make(chan *rpc.Call, 1))
	call = <-call.Done
	if call.Error != nil {
		fmt.Println("call  error: ", err.Error())
		return
	}

	fmt.Println(reply) // 56
}
