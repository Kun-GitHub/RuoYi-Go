package usecase

import (
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/pkg/cache"
	"go.uber.org/zap"
)

type AuthService struct {
	service input.SysUserService
	redis   *cache.RedisClient
	logger  *zap.Logger
}

func NewAuthService(service input.SysUserService, redis *cache.RedisClient, logger *zap.Logger) input.AuthService {
	return &AuthService{service: service, redis: redis, logger: logger}
}

func (this *AuthService) Login(l model.LoginRequest) (*model.LoginSuccess, error) {
	_, err := this.service.QueryUserByUserName(l.Username)

	return nil, err
}

func (this *AuthService) Logout(token string) error {
	if err := this.redis.Del(token); err != nil {
		this.logger.Debug("redis del error", zap.Error(err))
		return err
	}
	return nil
}
