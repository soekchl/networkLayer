/*
	创建者	:	Luke
	日期	:	2020年3月17日
	联系方式:	soekchl@163.com
*/
package main

import (
	. "github.com/soekchl/myUtils"
	nl "github.com/soekchl/networkLayer"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		Client()
	}()

	Server()
}

func Client() {
	client, err := nl.Dial("tcp", ":1111", 5)
	if err != nil {
		Error(err)
		return
	}
	Notice("客户端连接成功!")
	clientLoop(client)
}

func clientLoop(session *nl.Session) {
	data := &nl.FormatData{
		Id:   1,
		Seq:  2,
		Body: []byte{1, 3, 4, 5, 4},
	}
	err := session.Send(data)
	if err != nil {
		Error(err)
		return
	}
	Notice("Send Ok!")

	data, err = session.Receive()
	if err != nil {
		Error(err)
		return
	}

	Notice("Client Recv: ", data)
}

func Server() {
	server, err := nl.Listen("tcp", ":1111", 5, nl.HandlerFunc(serverLoop))
	if err != nil {
		Error(err)
		return
	}
	Notice("服务器开启！")
	server.Serve()
}

func serverLoop(session *nl.Session) {
	defer session.Close()
	Notice("服务器 接收连接：", session.RemoteAddr())
	for {
		data, err := session.Receive()
		if err != nil {
			Error(err)
			return
		}
		Notice(data)
		session.Send(data)
	}
}
