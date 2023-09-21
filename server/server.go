package server

import (
	"crypto/tls"
	"crypto/x509"
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

	// 使用系统时间
	h.Date = int(time.Now().Unix())
	hostData[h.Sid] = append(hostData[h.Sid], *h)
	if len(hostData[h.Sid]) > 90 {
		*result = " is much "
		base.HostData[h.Sid] = base.HostData[h.Sid][80:]
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

func TlsService() {

	// log.Println(key)
	//注册服务
	s := rpc.NewServer()
	s.Register(new(Server))

	cert, _ := tls.X509KeyPair(base.SCert, base.SKey)
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(base.CCert)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	l, err := tls.Listen("tcp", *base.Listen, config)
	fmt.Println("开始监听", *base.Listen)
	s.Accept(l)

	// https
	// hs := &http.Server{
	// 	Addr:           *base.Listen,
	// 	Handler:        s,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// 	TLSConfig:      &tls.Config{},
	// }
	// var err error
	// hs.TLSConfig.Certificates = make([]tls.Certificate, 1)

	// hs.TLSConfig.Certificates[0], err =
	// 	tls.X509KeyPair(base.Cert, base.Key)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println("开始监听", *base.Listen)
	// err = hs.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalln(err)
	}
}

func Service() {
	//注册服务
	s := rpc.NewServer()
	s.Register(new(Server))
	hs := &http.Server{
		Addr:           *base.Listen,
		Handler:        s,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//监听服务
	fmt.Println("开始监听", *base.Listen)
	err := hs.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
