package server

import (
	"context"
	"net"

	"github.com/samcsf/daily-go/pkg/blog/server/model"
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
			Title:      post.Title,
			Content:    post.Content,
			CreateAt:   post.Create_at.String(),
			ModifiedAt: post.Modified_at.String(),
		}
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

func (ps *postServer) CreatePost(ctx context.Context, post *pb.Post) (*pb.Empty, error) {
	p := &model.Post{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}

	service.Post.SavePost(p)
	return &pb.Empty{}, nil
}

func (ps *postServer) UpdatePost(ctx context.Context, post *pb.Post) (*pb.Empty, error) {
	p := &model.Post{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}

	service.Post.UpdatePost(p)
	return &pb.Empty{}, nil
}

func (ps *postServer) DeletePost(ctx context.Context, post *pb.Post) (*pb.Empty, error) {
	p := &model.Post{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}

	service.Post.DelPost(p)
	return &pb.Empty{}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":50051")
	util.ChkErr(err)
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &postServer{})
	err = s.Serve(lis)
	util.ChkErr(err)
}
