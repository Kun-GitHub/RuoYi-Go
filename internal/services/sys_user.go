package services

import (
	"RuoYi-Go/internal/models"
	rydb "RuoYi-Go/pkg/db"
	"encoding/json"
	"fmt"
)

func QueryUserByLoginName(loginname string, structEntity *models.SysUser) error {
	// 尝试从缓存中获取
	userBytes, err := rydb.DB.Cache.Get([]byte(fmt.Sprintf("LoginName:%s", loginname)))
	if err == nil {
		// 缓存命中
		json.Unmarshal(userBytes, &structEntity)
		return nil
	}

	err = rydb.DB.FindColumns(models.TableNameSysUser, structEntity,
		"login_name = ? and status = '0' and del_flag = '0'", loginname)
	if err != nil {
		return err
	} else {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			rydb.DB.Cache.Set([]byte(fmt.Sprintf("LoginName:%s", loginname)), userBytes, 0)        // 第三个参数是过期时间，0表示永不过期
			rydb.DB.Cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, 0) // 第三个参数是过期时间，0表示永不过期
		}
		return nil
	}
}
