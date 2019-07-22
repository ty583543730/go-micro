package main

import (
	hello "./proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
)

type Say struct {

}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rep *hello.Response) error {
	log.Println("Received Say.Hello request" + req.Name)
	rep.Msg = "Hello " + req.Name + " first go-micro"
	return nil
}

func main()  {
	service := micro.NewService(
		micro.Name("daniel.test.Hello"),
	)

	service.Init()

	hello.RegisterSayHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
