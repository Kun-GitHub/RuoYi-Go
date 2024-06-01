// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Shanghai",
		"postgres", "postgres", "postgresql.123456", "172.16.10.215",
		5432, "postgres")

	db, _ = gorm.Open(postgres.Open(dsn))
}

func main() {
	if db != nil {
		g := gen.NewGenerator(gen.Config{
			OutPath: "./",
		})

		g.UseDB(db)

		// generate all table from database
		g.ApplyBasic(g.GenerateAllTable()...)

		g.Execute()
	}
}
