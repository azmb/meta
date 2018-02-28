package main

import (
	"context"
	"log"
	"net"

	"github.com/azmb/meta/pb"

	"google.golang.org/grpc"
)

const bind = "127.0.0.1:9013"

type auth struct{}

func (auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{pb.TokenKey: pb.Token}, nil
}

func (auth) RequireTransportSecurity() bool { return false }

type server struct{}

func (*server) Do(ctx context.Context, t *pb.Test) (*pb.Test, error) {
	return &pb.Test{Name: "You got this!"}, nil
}

func main() {
	go serve()

	dial()
}

func serve() {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(pb.StreamInterceptor),
		grpc.UnaryInterceptor(pb.UnaryInterceptor),
	)

	pb.RegisterTestServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func dial() {
	var opts = []grpc.DialOption{
		grpc.WithPerRPCCredentials(&auth{}),
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(bind, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTestServiceClient(conn)

	resp, err := c.Do(context.Background(), &pb.Test{})
	if err != nil {
		log.Fatalf("RPC call failed: %s", err)
	}

	log.Printf("%#v", resp)
}
