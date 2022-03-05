package main

import (
	"github.com/FFFcomewhere/sk_object/pkg/bootstrap"
	conf "github.com/FFFcomewhere/sk_object/pkg/config"
	"github.com/FFFcomewhere/sk_object/pkg/mysql"
	"github.com/FFFcomewhere/sk_object/sk-admin/setup"
)

func main() {
	mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	//setup.InitEtcd()
	setup.InitZk()
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
}
