package grpcserver

import (
	"fww-core/internal/config"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Init(cfg *config.GrpcServerConfig) *grpc.Server {
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	return s
}

func StartService(s *grpc.Server, cfg *config.GrpcServerConfig, log *zap.SugaredLogger) error {
	lis, err := net.Listen("tcp", cfg.Host+":"+cfg.Port)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		return err
	}

	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
