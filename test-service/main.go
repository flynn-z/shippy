package main

import (
	"fmt"
	pb "shippy/test-service/proto/test"

	"github.com/labstack/gommon/log"
	"github.com/micro/go-micro"
)

func main() {
	// 连接到数据库
	db, err := CreateConnection()

	fmt.Printf("%+v\n", db)
	fmt.Printf("err: %v\n", err)

	defer db.Close()

	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	repo := &UserRepository{db}

	// 自动检查 User 结构是否变化
	db.AutoMigrate(&pb.User{})

	s := micro.NewService(
		micro.Name("go.micro.srv.test"),
		micro.Version("latest"),
	)

	s.Init()

	pb.RegisterTestServiceHandler(s.Server(), &handler{repo})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}

}
