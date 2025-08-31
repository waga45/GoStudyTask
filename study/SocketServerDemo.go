package study

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
)

var groupConnMap map[string]net.Conn

func StartSocketServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("服务器已启动")
	lock := sync.Mutex{}
	groupConnMap = make(map[string]net.Conn, 1024)
	for {
		conn, err := listener.Accept()
		if err != nil {
			listener.Close()
			panic(err)
			return
		}
		lock.Lock()
		groupConnMap[conn.RemoteAddr().String()] = conn
		lock.Unlock()
		sendBorodcast(conn.RemoteAddr().String(), "上线了，快来围观！")
		conn.Write([]byte("哇哦，星渊大陆等于等来你(" + conn.RemoteAddr().String() + ")\n"))
		go handlerConnect(conn)
	}
}

// 链接
func handlerConnect(conn net.Conn) {
	defer conn.Close()
	fmt.Println("有新链接进来了")
	buffer := make([]byte, 2048)
	currentAdd := conn.RemoteAddr().String()
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				sendBorodcast(currentAdd, "下线了，bye")
				fmt.Println("客户端", conn.RemoteAddr().String(), "链接断开")
				runtime.Goexit()
			}
			fmt.Println(err)
			return
		}
		cmd := string(buffer[:n])
		fmt.Println("服务器读取到数据:", cmd)
		sendBorodcast(currentAdd, cmd)

	}
}

func sendBorodcast(currentAddr string, content string) {
	if len(currentAddr) <= 0 {
		return
	}
	content = fmt.Sprintf("%s:%s", currentAddr, content)
	for k, v := range groupConnMap {
		if k != currentAddr {
			v.Write([]byte(content + "\n"))
		}
	}
}
