package main

import (
	loggerx "github.com/b3liv3r/logger"
	userv1 "github.com/b3liv3r/protos-for-gym/gen/go/user"
	"github.com/b3liv3r/users-for-gym/config"
	"github.com/b3liv3r/users-for-gym/modules/db"
	"github.com/b3liv3r/users-for-gym/modules/user/repository"
	"github.com/b3liv3r/users-for-gym/modules/user/service"
	server "github.com/b3liv3r/users-for-gym/modules/user/urpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	appConf := config.MustLoadConfig()

	logger := loggerx.InitLogger(appConf.Name, appConf.Production)

	sqlDB, err := db.NewSqlDB(logger, appConf.Db)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	repo := repository.NewUserRepositoryDB(sqlDB)
	service := service.NewUserService(repo, logger)
	s := InitRPC(service)
	lis, err := net.Listen("tcp", appConf.GrpcServerPort)
	if err != nil {
		logger.Error("failed to listen:", zap.Error(err))
	}
	logger.Info("grpc server listening at", zap.Stringer("address", lis.Addr()))
	if err = s.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", zap.Error(err))
	}
}

func InitRPC(uservice service.Userer) *grpc.Server {
	s := grpc.NewServer()
	userv1.RegisterUserServer(s, server.NewUserRPCServer(uservice))

	return s
}
