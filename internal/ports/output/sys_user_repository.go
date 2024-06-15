package output

import (
	"RuoYi-Go/internal/domain/model"
)

type SysUserRepository interface {
	QueryUserByUserName(username string) (*model.SysUser, error)
}
