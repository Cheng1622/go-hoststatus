package base

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"sync"
	"time"
)

type HostInfo struct {
	Sid      string
	HostName string
	SysInfo  string
	Ip       string
	Sip      string
	Mem      string
	Cpu      string
	Disk     string
	Date     int
}

func (t *HostInfo) Bytes() []byte {
	a := bytes.NewBuffer(nil)
	fmt.Fprintln(a, "Sid", t.Sid, "				</br>")
	fmt.Fprintln(a, "HostName", t.HostName, "				</br>")
	fmt.Fprintln(a, "SysInfo", t.SysInfo, "				</br>")
	fmt.Fprintln(a, "Ip", t.Ip, "				</br>")
	fmt.Fprintln(a, "Mem", t.Mem, "				</br>")
	fmt.Fprintln(a, "Cpu", t.Cpu, "				</br>")
	fmt.Fprintln(a, "Disk", t.Disk, "				</br>")
	d := time.Unix(int64(t.Date), 0).Local().Format("01/02 15:04:05")
	fmt.Fprintln(a, "Date", d, "				</br>")
	return a.Bytes()
}
func (t *HostInfo) String() string {
	return string(t.Bytes())
}

var HostData = make(map[string][]HostInfo)
var HostDataLock = new(sync.RWMutex)
var (
	Is_server *bool
	Is_user   *bool
	Listen    *string
	LosTime   *int
)

func init() {
	Is_server = flag.Bool("s", false, "server")
	Is_user = flag.Bool("u", false, "getdata")
	Listen = flag.String("l", ":12345", "listen addr")
	LosTime = flag.Int("t", 60, "Lost Time for alert /s")
	flag.Parse()
}

// # 生成私钥
// openssl genrsa -out server.key 2048
// # 生成证书
// openssl req -new -x509 -key server.key -out server.crt -days 3650
// # 只读权限
// chmod 400 server.key
// openssl genrsa -out server.key 2048 &&openssl req -new -x509 -key server.key -out server.crt -days 3650
// openssl genrsa -out client.key 2048 &&openssl req -new -x509 -key client.key -out client.crt -days 3650

// //go:embed pem/fullchain.pem
// var Cert []byte

// //go:embed pem/privkey.pem
// var Key []byte

//go:embed pem/client.crt
var CCert []byte

//go:embed pem/client.key
var CKey []byte

//go:embed pem/server.crt
var SCert []byte

//go:embed pem/server.key
var SKey []byte
