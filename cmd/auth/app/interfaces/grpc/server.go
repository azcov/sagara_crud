package grpc

import (
	"context"

	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
	"github.com/azcov/sagara_crud/cmd/auth/proto"

	grpcInterceptor "github.com/azcov/sagara_crud/internal/grpc/server/interceptor"

	"google.golang.org/protobuf/types/known/emptypb"
)

type authServer struct {
	proto.UnimplementedAuthServiceServer
	usecase domain.AuthenticationUsecase
}

func NewServer(u domain.AuthenticationUsecase) proto.AuthServiceServer {
	return &authServer{
		usecase: u,
	}
}

func (s authServer) ValidationToken(ctx context.Context, _ *emptypb.Empty) (*proto.ValidationTokenResponse, error) {
	res := &proto.ValidationTokenResponse{
		UserId: ctx.Value(grpcInterceptor.MetadataKeyUserID).(string),
		RoleId: ctx.Value(grpcInterceptor.MetadataKeyRoleID).(int32),
	}
	return res, nil
}
