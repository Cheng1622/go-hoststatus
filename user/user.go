package user

import (
	"fmt"
	"log"
	"net/rpc"
	"sort"

	"github.com/Cheng1622/go-hoststatus/base"
)

func showHostData() {
	fmt.Printf("\x1b[2J")
	ks := make([]string, 0, len(base.HostData))
	for k := range base.HostData {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		v := base.HostData[k]
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
