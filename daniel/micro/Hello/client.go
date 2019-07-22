package main

import (
	hello "./proto"
	common "./common"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"log"
)

func main()  {
	config.LoadFile("./config/db.json")
	var cfg common.GlobalCfg
	config.Scan(&cfg)

	fmt.Println(cfg.Consul["host"])
	fmt.Println(cfg.Mysql["xproject"])
	fmt.Println(cfg.Mysql["renewal"])
	fmt.Println(cfg.Mysql["deal"])

 	service := micro.NewService(micro.Name("daniel"))

	service.Init()

	client := hello.NewSayService("daniel.test.Hello", service.Client())

	rsp, err := client.Hello(context.TODO(), &hello.Request{Name:"daniel"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp.GetMsg())
}
