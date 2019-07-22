package main

import (
	proto "../../proto/renewal"
	renewalDao "../../dao/renewal"
	renewalHandle "../../handle/renewal"
	"github.com/micro/go-micro"
	"log"
)

func main()  {
	// 实例化dao
	renewalDealDao, err := renewalDao.NewRenewalDealDao()
	if err != nil {
		log.Fatal(err) // todo 写日志文件
	}

	// 这里可以封装起来,只提供一个实例化service的接口
	renewalReadHandle := new(renewalHandle.RenewalDealReadAoHandle)
	renewalReadHandle.RenewalDealDao = *renewalDealDao

	service := micro.NewService(
		micro.Name("go.micro.srv.getRenewalDealByDealId"),
	)

	service.Init()

	proto.RegisterRenewalDealReadAoServiceHandler(service.Server(), renewalReadHandle)

	if err := service.Run(); err != nil {
		log.Fatal(err)  // todo 写日志文件
	}
}
