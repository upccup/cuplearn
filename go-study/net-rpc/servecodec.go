package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"net/rpc"
	"time"
)

// func ServeCodec(codec ServerCodec)
// 功能说明: 这个函数运行一个但连接的DefaultServer, 服务一直开启,知道客户端断开连接, 并且可以设置ServeCodec,
// 			用于处理请求和回复. 实际上就是调用DefaultServer.ServeCodec(codec)这个方法

type Args struct {
	A, B int
}

type Quptient struct {
	Qup, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quptient) error {
	if args.B == 0 {
		return fmt.Errorf("divide by zero")
	}
	quo.Qup = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

type myServerCodec struct {
	rwc    io.ReadWriteCloser
	dec    *gob.Decoder
	enc    *gob.Encoder
	encBuf *bufio.Writer
}

func (c *myServerCodec) ReadRequestHeader(r *rpc.Request) error {
	fmt.Println("调用: ReadRequestHeader ")
	return c.dec.Decode(r)
}

func (c *myServerCodec) ReadRequestBody(body interface{}) error {
	fmt.Println("调用: ReadRequestBody ")
	return c.dec.Decode(body)
}

func (c *myServerCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {
	fmt.Println("调用: WriteResponse  ")
	if err = c.enc.Encode(r); err != nil {
		return
	}

	if err = c.enc.Encode(body); err != nil {
		return
	}

	return c.encBuf.Flush()
}

func (c *myServerCodec) Close() error {
	fmt.Println("调用: Close ")
	return c.rwc.Close()
}

func myAccept(l net.Listener) {
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Accept error: ", err)
		return
	}

	// 创建一个自定义的ServerCode, 实际上这些代码就从go源码中获取出来的
	buf := bufio.NewWriter(conn)
	codec := &myServerCodec{conn, gob.NewDecoder(conn), gob.NewEncoder(buf), buf}

	//运行连接
	go rpc.ServeCodec(codec)
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
		fmt.Println("listen error: ", err)
		return
	}

	// 运行服务监听
	go myAccept(l)

	// 暂停2秒, 让服务有足够的时间开启
	time.Sleep(time.Second * 2)

	// 设置客户端
	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("ResolveTCPAddr has error: ", err)
		return
	}

	// 使用TCP的方式进行网络请求
	conn, _ := net.DialTCP("tcp", nil, address)
	defer conn.Close()
	// 创建一个客户端
	client := rpc.NewClient(conn)
	defer client.Close()

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("arith error: ", err)
		return
	}

	fmt.Println(reply)

	args2 := &Args{17, 8}
	var reply2 Quptient
	err = client.Call("Arith.Divide", args2, &reply2)
	if err != nil {
		fmt.Println("arith error: ", err)
		return
	}

	fmt.Println(reply2)
}

// 结果
/*
调用: ReadRequestHeader
调用: ReadRequestBody
调用: ReadRequestHeader
调用: WriteResponse
56
调用: ReadRequestBody
调用: ReadRequestHeader
调用: WriteResponse
{2 1}

*/
