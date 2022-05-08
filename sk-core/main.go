package main

import (
	"github.com/FFFcomewhere/seckill/sk-core/setup"
)

func main() {
	setup.InitZk()
	setup.InitRedis()
	setup.RunService()
}
