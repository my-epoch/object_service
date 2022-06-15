package handler

import (
	"context"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetNearest(_ context.Context, _ *pb.GetNearestRequest) (*pb.GetNearestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearest not implemented")
}
