package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/pkg/cache"
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type SysUserOnlineService struct {
	redis  *cache.RedisClient
	logger *zap.Logger
}

func NewSysUserOnlineService(redis *cache.RedisClient, logger *zap.Logger) input.SysUserOnlineService {
	return &SysUserOnlineService{redis: redis, logger: logger}
}

func (s *SysUserOnlineService) List(page, size int, ipaddr, userName string) ([]*model.SysUserOnline, int64, error) {
	keys, err := s.redis.Keys(fmt.Sprintf("%s:*", common.TOKEN))
	if err != nil {
		return nil, 0, err
	}

	var userList []*model.SysUserOnline
	for _, key := range keys {
		val, err := s.redis.Get(key)
		if err != nil {
			continue
		}

		var userOnline model.SysUserOnline
		if err := json.Unmarshal([]byte(val), &userOnline); err != nil {
			continue
		}

		if ipaddr != "" && !strings.Contains(userOnline.Ipaddr, ipaddr) {
			continue
		}
		if userName != "" && !strings.Contains(userOnline.UserName, userName) {
			continue
		}

		userList = append(userList, &userOnline)
	}

	// Pagination
	total := int64(len(userList))
	start := (page - 1) * size
	end := start + size
	if start > int(total) {
		start = int(total)
	}
	if end > int(total) {
		end = int(total)
	}

	return userList[start:end], total, nil
}

func (s *SysUserOnlineService) ForceLogout(tokenId string) error {
	return s.redis.Del(fmt.Sprintf("%s:%s", common.TOKEN, tokenId))
}
