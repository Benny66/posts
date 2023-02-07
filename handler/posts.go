package handler

import (
	"context"

	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"

	pb "github.com/Benny66/posts/proto"
)

type Posts struct{}

// Return a new handler
func New() *Posts {
	return &Posts{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Posts) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Info("Received Posts.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Posts) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.Posts_StreamStream) error {
	log.Infof("Received Posts.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Posts) PingPong(ctx context.Context, stream pb.Posts_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func (p *Posts) Save(ctx context.Context, req *pb.SaveRequest, rsp *pb.SaveResponse) error {
	if req.Post == nil {
		return errors.BadRequest("posts.Save", "req nil")
	}
	if len(req.Post.Id) == 0 || len(req.Post.Title) == 0 || len(req.Post.Content) == 0 {
		return errors.BadRequest("posts.Save", "ID, title or content is missing")
	}
	return nil
}
