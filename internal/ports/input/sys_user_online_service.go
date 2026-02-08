package input

import "RuoYi-Go/internal/domain/model"

type SysUserOnlineService interface {
	List(page, size int, ipaddr, userName string) ([]*model.SysUserOnline, int64, error)
	ForceLogout(tokenId string) error
}
