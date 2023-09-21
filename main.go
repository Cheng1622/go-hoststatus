package main

import (
	"github.com/Cheng1622/go-hoststatus/base"
	"github.com/Cheng1622/go-hoststatus/client"
	"github.com/Cheng1622/go-hoststatus/server"
	"github.com/Cheng1622/go-hoststatus/user"
)

func main() {
	if *base.Is_server {
		server.Service()
		return
	}

	if *base.Is_user {
		user.User()
		return
	}

	// 客户端
	// t := time.NewTicker(time.Minute / 10)
	// defer t.Stop()
	// for {
	client.Client()
	// 	<-t.C
	// }

}
