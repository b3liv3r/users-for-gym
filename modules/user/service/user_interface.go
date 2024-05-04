package service

import (
	"context"
	"github.com/b3liv3r/users-for-gym/modules/user/models"
)

type Userer interface {
	Create(ctx context.Context, userId int, username, phone, email string) (string, error)
	Profile(ctx context.Context, userId int) (models.User, error)
	Update(ctx context.Context, userId int, username, phone, email string) (string, error)
	UpdateSubLvl(ctx context.Context, userId int, level int) (string, error)
	UpdateGymID(ctx context.Context, userId int, gymId int) (string, error)
}
