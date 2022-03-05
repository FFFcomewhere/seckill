package setup

import (
	conf "github.com/FFFcomewhere/sk_object/pkg/config"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

//初始化Etcd
func InitEtcd() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"39.98.179.73:2379"}, // conf.Etcd.Host
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}
	conf.Etcd.EtcdSecProductKey = "product"
	conf.Etcd.EtcdConn = cli
}
