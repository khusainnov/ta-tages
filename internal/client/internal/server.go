package internal

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConn() (*grpc.ClientConn, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(":9001", opts)
	if err != nil {
		return nil, fmt.Errorf("error due dial with client, %w", err)
	}

	return conn, err
}
