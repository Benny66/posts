package main

import (
	"github.com/Benny66/posts/handler"
	pb "github.com/Benny66/posts/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("posts"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPostsHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
