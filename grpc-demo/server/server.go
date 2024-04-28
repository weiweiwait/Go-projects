package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

// ServiceA 自定义一个结构体类型
type ServiceA struct {
}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

//	func main() {
//		service := new(ServiceA)
//		rpc.Register(service)
//		rpc.HandleHTTP() //基于Http写一
//		l, e := net.Listen("tcp", ":9091")
//		if e != nil {
//			log.Fatal("listen error:", e)
//		}
//		http.Serve(l, nil)
//	}
//
// 基于TCP
//
//	func main() {
//		service := new(ServiceA)
//		rpc.Register(service) // 注册RPC服务
//		l, e := net.Listen("tcp", ":9091")
//		if e != nil {
//			log.Fatal("listen error:", e)
//		}
//		for {
//			conn, _ := l.Accept()
//			rpc.ServeConn(conn)
//		}
//	}
//
// 使用JSON写一
func main() {
	service := new(ServiceA)
	rpc.Register(service) // 注册RPC服务
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, _ := l.Accept()
		// 使用JSON协议
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
