package main

import (
	"github.com/FFFcomewhere/seckill/pkg/bootstrap"
	conf "github.com/FFFcomewhere/seckill/pkg/config"
	"github.com/FFFcomewhere/seckill/pkg/mysql"
	"github.com/FFFcomewhere/seckill/sk-admin/setup"
)

func main() {
	mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	//setup.InitEtcd()
	setup.InitZk()
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
}
