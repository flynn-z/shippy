package main

import (
	"context"
	"fmt"
	"log"
	pb "shippy/test-service/proto/test"
)

type handler struct {
	repo Repository
}

func (h *handler) Create(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	log.Println("Creating user: ", &req)
	fmt.Println("user : ", req.User)
	if err := h.repo.Create(req.User); err != nil {
		return nil
	}
	resp.User = req.User
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	u, err := h.repo.Get(req.User.Id)
	if err != nil {
		return err
	}
	resp.User = u
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	resp.Users = users
	return nil
}
