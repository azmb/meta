package pb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	Token    = "important\n"
	TokenKey = "token"
)

func StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := authorize(stream.Context()); err != nil {
		return err
	}

	return handler(srv, stream)
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := authorize(ctx); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func authorize(ctx context.Context) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md[TokenKey]) != 1 {
			return status.Error(codes.Unauthenticated, "expected single token in request metadata")
		}

		if md[TokenKey][0] != Token {
			return status.Error(codes.PermissionDenied, "token unknown")
		} else {
			return nil
		}
	}

	return status.Error(codes.Unauthenticated, "request metadata not present")
}
