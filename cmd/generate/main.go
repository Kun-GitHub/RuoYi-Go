// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

var db *gorm.DB

func init() {
	//dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Shanghai",
	//	"postgres", "postgres", "postgresql.123456", "172.16.10.215",
	//	5432, "postgres")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root",
		"hn02le.34lkdLKD", "192.168.3.24", 3306, "ruoyi")
	db, _ = gorm.Open(mysql.Open(dsn))
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if i == 0 {
			parts[i] = strings.ToLower(parts[i])
		} else {
			parts[i] = strings.Title(parts[i])
		}
	}
	return strings.Join(parts, "")
}

func main() {
	if db != nil {
		config := gen.Config{
			OutPath:      "./dao",
			ModelPkgPath: "./model",
			Mode:         gen.WithDefaultQuery,
		}
		config.WithJSONTagNameStrategy(func(columnName string) string {
			return toCamelCase(columnName)
		})

		g := gen.NewGenerator(config)

		g.UseDB(db)

		// generate all table from database
		g.ApplyBasic(g.GenerateAllTable()...)

		g.Execute()
	}
}
