package main

import (
	"./common"
	"./logic"
	proto "./proto"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

func main() {
	// 加载配置
	var cfg common.GlobalCfg
	config.LoadFile("./config/db.json")
	config.Scan(&cfg)

	fmt.Println(cfg)

	reg := consul.NewRegistry(func(ops *registry.Options) {
		ops.Addrs = []string{
			cfg.Consul["host"],
		}
	})

	// 定义服务
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("renewalDealAo"),
	)

	// 实例化handle
	renewalLogic, err := logic.NewRenewalReadAoLogic(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化
	service.Init()

	// 注册服务handle
	proto.RegisterRenewalReadAoHandler(service.Server(), renewalLogic)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
