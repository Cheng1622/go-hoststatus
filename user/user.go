package user

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/Cheng1622/go-hoststatus/base"
)

func showHostData() {
	for _, v := range base.HostData {
		v1 := v[len(v)-1]
		fmt.Println(v1.String())
	}
}
func User() {
	//连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", *base.Listen)
	if err != nil {
		log.Println(err)
	}

	//调用方法
	result := base.HostData
	err = conn.Call("Server.GetData", 1, &result)
	showHostData()

	if err != nil {
		log.Println(err)
		return
	}
}
