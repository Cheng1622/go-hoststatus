package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"time"

	"github.com/Cheng1622/go-hoststatus/base"
)

var hostData = make(map[string][]HostInfo)

type HostInfo struct {
	Sid      string
	HostName string
	SysInfo  string
	Ip       string
	Mem      string
	Cpu      string
	Disk     string
	Date     int
}
type Server struct {
}

func (l *Server) GetData(h int, result *map[string][]HostInfo) error {
	*result = hostData
	return nil
}

func liten() {
	t := time.NewTicker(time.Minute)
	for ; ; base.HostDataLock.Unlock() {
		<-t.C
		base.HostDataLock.Lock()
		if len(base.HostData) == 0 {
			continue
		}
		for _, v := range base.HostData {
			// last push time
			h := v[len(v)-1]
			t := int(time.Now().Unix()) - h.Date
			if t > 60 {
				// alert
				// base.Mail.Set(base.UserMail, "host lost "+h.Sid, h.String()).Send()
				delete(base.HostData, h.Sid)
			}
		}
	}
}

func (l *Server) Save(h *HostInfo, result *string) error {
	*result = "I see"
	log.Println("recive a msg")
	if h.Sid == "" {
		return errors.New("sid is null")
	}
	base.HostDataLock.Lock()
	defer base.HostDataLock.Unlock()

	if hostData[h.Sid] == nil {
		hostData[h.Sid] = make([]HostInfo, 0)
		log.Println("find a new host")
		// base.Mail.Set(base.UserMail, "HostListen find a new host", h.String()).Send()
	}

	hostData[h.Sid] = append(hostData[h.Sid], *h)
	if len(hostData[h.Sid]) > 90 {
		*result = " is much "
		// 转储
	}
	// fmt.Println()
	// fmt.Println("Sid", h.Sid)
	// fmt.Println("HostName", h.HostName)
	// fmt.Println("SysInfo", h.SysInfo)
	// fmt.Println("Ip", h.Ip)
	// fmt.Println("Mem", h.Mem)
	// fmt.Println("Cpu", h.Cpu)
	// fmt.Println("Disk", h.Disk)
	// fmt.Println("Date", h.Date)
	return nil
}

func init() {
	// 开启监听，失联报警
	go liten()
}

func Service() {
	//注册服务
	rpc.Register(new(Server))
	//绑定http协议
	rpc.HandleHTTP()
	//监听服务
	fmt.Println("开始监听", *base.Listen)
	err := http.ListenAndServe(*base.Listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
