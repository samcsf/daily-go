package server

import (
	"net"

	"github.com/samcsf/daily-go/pkg/blog/server/service"
	"github.com/samcsf/daily-go/pkg/util"
	"google.golang.org/grpc"

	pb "github.com/samcsf/daily-go/pkg/blog/server/grpc/post/post"
)

type postServer struct{}

func (ps *postServer) GetPosts(em *pb.Empty, stream pb.PostService_GetPostsServer) error {
	posts, err := service.Post.GetPosts()
	util.ChkErr(err)

	for _, post := range posts {
		p := &pb.Post{
			Title:    post.Title,
			Content:  post.Content,
			CreateAt: post.Create_at.String(),
		}
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":50051")
	util.ChkErr(err)
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &postServer{})
	err = s.Serve(lis)
	util.ChkErr(err)
}
