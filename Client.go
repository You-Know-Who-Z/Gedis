package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// TCP 客户端
func Client(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("client err: %v", err))
	}
	// 关闭TCP连接
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户输入
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			// 如果输入q就退出
			return
		}
		// 发送数据
		go send(inputInfo, conn)
	}
}
func send(input string, conn net.Conn) {
	_, err := conn.Write([]byte(input + "\n"))
	if err != nil {
		log.Fatal("send failed, err:", err)
		return
	}
	log.Println("ss")
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		log.Fatal("recv failed, err:", err)
		return
	}
	fmt.Println(string(buf[:n]))
}
