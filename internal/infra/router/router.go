package router

import (
	"context"
	"log"
	"net"
	mediav1 "neuro-most/media-service/gen/go/media/v1"
	"neuro-most/media-service/internal/adapters/api/action"
	"neuro-most/media-service/internal/adapters/presenter"
	"neuro-most/media-service/internal/adapters/repo"
	"neuro-most/media-service/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Router struct {
	db repo.GSQL
	mediav1.UnimplementedMediaServiceServer
}

func NewRouter(db repo.GSQL) Router {
	return Router{
		db: db,
	}
}

func (r *Router) Listen() {
	port := ":3001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts = []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	mediav1.RegisterMediaServiceServer(srv, r)

	log.Printf("Starting gRPC server on port %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (r *Router) CreateMedia(ctx context.Context, input *mediav1.CreateMediaRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewCreateMediaUseCase(repo.NewMediaRepo(r.db))
		act = action.NewCreateMediaAction(uc)
	)

	err := act.Execute(ctx, input)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (r *Router) DeleteMedia(ctx context.Context, input *mediav1.DeleteMediaRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewDeleteMediaUseCase(repo.NewMediaRepo(r.db))
		act = action.NewDeleteMediaAction(uc)
	)
	err := act.Execute(ctx, input)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (r *Router) GetMediaById(ctx context.Context, input *mediav1.GetMediaByIdRequest) (*mediav1.Media, error) {
	var (
		uc  = usecase.NewGetByIDMediaUseCase(repo.NewMediaRepo(r.db), presenter.NewGetByIDMediaPresenter())
		act = action.NewGetByIDMediaAction(uc)
	)

	output, err := act.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *Router) GetMediaFeed(ctx context.Context, input *mediav1.GetMediaFeedRequest) (*mediav1.GetMediaFeedResponse, error) {
	var (
		uc  = usecase.NewGetAllMediaUseCase(repo.NewMediaRepo(r.db), presenter.NewGetAllMediaPresenter())
		act = action.NewGetAllMediaAction(uc)
	)

	output, err := act.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *Router) UpdateMedia(ctx context.Context, input *mediav1.UpdateMediaRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewUpdateMediaUseCase(repo.NewMediaRepo(r.db))
		act = action.NewUpdateMediaAction(uc)
	)

	err := act.Execute(ctx, input)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
