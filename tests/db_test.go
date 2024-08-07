// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package main

import (
	"gorm.io/gorm"
	"testing"
)

// 定义模型（示例为User模型）
type User struct {
	gorm.Model
	Name  string
	Email string
}

func TestSqlite(t *testing.T) {

	//// 创建sqliteStruct实例
	//sqliteService := &dao.DatabaseStruct{}
	//
	//// 打开数据库并执行自动迁移（假设User结构体是你的模型）
	//err := sqliteService.OpenSqlite()
	//if err != nil {
	//	fmt.Println("Failed to open or migrate the database:", err)
	//	return
	//}
	//
	//// 在这里执行数据库相关操作，如查询、插入、更新等
	//// ...
	////sqliteService.Create(&User{Name: "张三", Email: "zhangsan@example.com"})
	//
	//// 完成所有操作后，关闭数据库连接
	//defer func() {
	//	err = sqliteService.CloseDB()
	//	if err != nil {
	//		fmt.Println("Failed to close the database connection:", err)
	//	}
	//}()

}
