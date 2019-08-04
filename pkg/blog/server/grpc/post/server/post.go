package server

import (
	"context"
	"log"
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

func (ps *postServer) CreatePost(ctx context.Context, post *pb.Post) (*pb.ExecResult, error) {
	p := &model.Post{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}

	res, err := service.Post.SavePost(p)
	if err != nil {
		log.Println("Grpc error when calling SavePost(): %v", err)
	}
	lid, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()
	return &pb.ExecResult{
		LastInsertId: lid,
		RowsAffected: rows,
	}, nil
}

func (ps *postServer) UpdatePost(ctx context.Context, post *pb.Post) (*pb.ExecResult, error) {
	p := &model.Post{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}

	res, err := service.Post.UpdatePost(p)
	if err != nil {
		log.Println("Grpc error when calling UpdatePost(): %v", err)
	}
	lid, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()
	return &pb.ExecResult{
		LastInsertId: lid,
		RowsAffected: rows,
	}, nil

}

func (ps *postServer) DeletePost(ctx context.Context, post *pb.Post) (*pb.ExecResult, error) {
	p := &model.Post{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}

	res, err := service.Post.DelPost(p)
	if err != nil {
		log.Println("Grpc error when calling DelPost(): %v", err)
	}
	lid, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()
	return &pb.ExecResult{
		LastInsertId: lid,
		RowsAffected: rows,
	}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":50051")
	util.ChkErr(err)
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &postServer{})
	log.Println("gRPC server start listening on tcp:50051")
	err = s.Serve(lis)
	util.ChkErr(err)
}
