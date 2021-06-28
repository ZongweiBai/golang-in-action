package core

import (
	"bufio"
	"io"
	"net"
	"strconv"

	"github.com/ZongweiBai/golang-in-action/config"
	"github.com/ZongweiBai/golang-in-action/core/proto"
)

func InitSocketServer() {
	listen, err := net.Listen("tcp", config.CONFIG.Socket.Host+":"+strconv.Itoa(config.CONFIG.Socket.Port))
	if err != nil {
		config.LOG.Errorf("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			config.LOG.Errorf("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	reader := bufio.NewReader(conn)
	for {
		recvStr, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			config.LOG.Errorf("decode msg failed", err)
			return
		}
		// var buf [128]byte
		// n, err := reader.Read(buf[:]) // 读取数据
		// if err != nil {
		// 	config.LOG.Errorf("read from client failed, err:", err)
		// 	break
		// }
		// recvStr := string(buf[:n])
		config.LOG.Infof("收到client端发来的数据：%s", recvStr)
		// conn.Write([]byte(recvStr)) // 发送数据
		data, err := proto.Encode(recvStr)
		if err != nil {
			config.LOG.Errorf("encode msg failed", err)
			return
		}
		conn.Write(data)
	}
}
