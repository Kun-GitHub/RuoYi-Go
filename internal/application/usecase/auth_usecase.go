package usecase

import (
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"go.uber.org/zap"
)

type AuthService struct {
	repo   output.SysUserRepository
	redis  *cache.RedisClient
	logger *zap.Logger
}

func NewAuthService(repo output.SysUserRepository, redis *cache.RedisClient, logger *zap.Logger) input.AuthService {
	return &AuthService{repo: repo, redis: redis, logger: logger}
}

func (this *AuthService) Login(l model.LoginRequest) (*model.LoginSuccess, error) {
	_, err := this.repo.QueryUserByUserName(l.Username)

	return nil, err
}

func (this *AuthService) Logout(token string) error {
	return this.redis.Del(token)
}
