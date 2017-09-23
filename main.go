package main

import (
	"Airttp/modules"
	"net/rpc"
	"net"
	"log"
	"fmt"
	"strconv"
)

type Http int

func main() {
	http := new(Http)

	server := rpc.NewServer()
	server.RegisterName("Http", http)

	l, e := net.Listen("tcp", ":5004")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("server start")
	server.Accept(l)
}

func (t *Http) Module(params modules.ModuleParams, result *modules.ModuleParams) error {
	fmt.Print("New Request : ")
	result.Copy(params)
	result.Res.Raw = []byte("HTTP/1.1 " + strconv.Itoa(result.Res.Code) + " " + result.Res.Message + "\r\n")
	for key, value := range result.Res.Headers {
		result.Res.Raw = append(result.Res.Raw[:], []byte(key + ": " + value + "\r\n")[:]...)
	}
	result.Res.Raw = append(result.Res.Raw[:], []byte("\r\n")[:]...)
	result.Res.Raw = append(result.Res.Raw[:], result.Res.Body[:]...)
	fmt.Println("OK")
	return nil
}