package server

import (
	"fmt"
	"net"

	"github.com/ondbyte/cloud_bees/blog"
	"google.golang.org/grpc"
)

// blocks
func Run(port uint) error {
	grpcServer := grpc.NewServer()
	blog.RegisterBlogServiceServer(grpcServer, NewBlogServiceServer())
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return grpcServer.Serve(listener)
}
