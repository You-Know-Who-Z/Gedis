package main

import "strings"

// 存储
var KvStore = make(map[string]string)

func set(k string, v string) {
	println("存储数据")
	KvStore[k] = v
}

func get(k string) string {
	println("获取数据")
	return KvStore[k]
}

// 命令行

func Command(comm string) string {
	d := strings.Split(comm, " ")
	println(d[0])
	switch d[0] {
	case "set":
		set(d[1], d[2])
	case "get":
		println(get(d[1]))
	}
	return "1"
}
