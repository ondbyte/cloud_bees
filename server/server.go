package server

import (
	"github.com/ondbyte/cloud_bees/blog"
	"google.golang.org/grpc"
)

func Run(port uint) {
	grpcServer := grpc.NewServer()
	blog.RegisterBlogServiceServer(grpcServer, &BlogServiceServer{})
}
