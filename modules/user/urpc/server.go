package server

import (
	"context"
	userv1 "github.com/b3liv3r/protos-for-gym/gen/go/user"
	"github.com/b3liv3r/users-for-gym/modules/user/service"
)

type UserRPCServer struct {
	userv1.UnimplementedUserServer
	srv service.Userer
}

func NewUserRPCServer(srv service.Userer) userv1.UserServer {
	return &UserRPCServer{srv: srv}
}

func (s *UserRPCServer) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	message, err := s.srv.Create(ctx, int(req.UserId), req.Username, req.Phone, req.Email)
	if err != nil {
		return nil, err
	}
	return &userv1.CreateResponse{Message: message}, nil
}

func (s *UserRPCServer) Profile(ctx context.Context, req *userv1.ProfileRequest) (*userv1.ProfileResponse, error) {
	user, err := s.srv.Profile(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &userv1.ProfileResponse{
		Username:        user.Username,
		Phone:           user.Phone,
		Email:           user.Email,
		SubscriptionLvl: int64(user.SubscriptionLvl),
		CurrentGymId:    int64(user.CurrentGymId),
	}, nil
}

func (s *UserRPCServer) Update(ctx context.Context, req *userv1.UpdateRequest) (*userv1.UpdateResponse, error) {
	message, err := s.srv.Update(ctx, int(req.UserId), req.Username, req.Phone, req.Email)
	if err != nil {
		return nil, err
	}
	return &userv1.UpdateResponse{Message: message}, nil
}

func (s *UserRPCServer) ChangeSubscriptions(ctx context.Context, req *userv1.ChangeSubscriptionsRequest) (*userv1.ChangeSubscriptionsResponse, error) {
	message, err := s.srv.UpdateSubLvl(ctx, int(req.UserId), int(req.SubscriptionLvl))
	if err != nil {
		return nil, err
	}
	return &userv1.ChangeSubscriptionsResponse{Message: message}, nil
}

func (s *UserRPCServer) ChangeCurrentGym(ctx context.Context, req *userv1.ChangeCurrentGymRequest) (*userv1.ChangeCurrentGymResponse, error) {
	message, err := s.srv.UpdateGymID(ctx, int(req.UserId), int(req.CurrentGymId))
	if err != nil {
		return nil, err
	}
	return &userv1.ChangeCurrentGymResponse{Message: message}, nil
}
