package setup

import (
	"fmt"
	conf "github.com/FFFcomewhere/seckill/pkg/config"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

//初始化Etcd	log.Printf("Connect zk sucess %s", conf.Zk.SecProductKey)
func InitZk() {
	var hosts = []string{"zk1:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	conf.Zk.ZkConn = conn
	conf.Zk.SecProductKey = "/product"
}
