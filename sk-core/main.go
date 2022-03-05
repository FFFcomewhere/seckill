package main

import (
	"github.com/FFFcomewhere/sk_object/sk-core/setup"
)

func main() {
	setup.InitZk()
	setup.InitRedis()
	setup.RunService()
}
