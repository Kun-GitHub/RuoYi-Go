// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package rydb

import (
	"RuoYi-Go/pkg/config"
	"database/sql"
	"fmt"
	"github.com/coocood/freecache"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
	"time"
)

type DatabaseStruct struct {
	db    *gorm.DB
	sqlDB *sql.DB
	cache *freecache.Cache

	mu sync.Mutex
}

var DB *DatabaseStruct

func (ds *DatabaseStruct) OpenSqlite() error {
	var err error = nil
	var dialector gorm.Dialector = nil

	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Shanghai",
		config.Conf.Database.DBtype, config.Conf.Database.User, config.Conf.Database.Password, config.Conf.Database.Host,
		config.Conf.Database.Port, config.Conf.Database.DBName)

	switch config.Conf.Database.DBtype {
	case "postgresql":
		dialector = postgres.Open(dsn)
	case "sqlite":
		dsn = "./test.db"
		dialector = sqlite.Open(dsn)
	default:
		dialector = mysql.Open(dsn)
	}
	ds.db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	ds.sqlDB, err = ds.db.DB()
	if err != nil {
		return err
	}

	// 设置连接池参数（例如设置最大空闲连接数和最大打开连接数）
	ds.sqlDB.SetMaxIdleConns(10)  // 设置空闲连接池中的最大连接数
	ds.sqlDB.SetMaxOpenConns(100) // 设置数据库的最大打开连接数

	// 如果需要，还可以设置连接在空闲多久后关闭（单位秒）
	ds.sqlDB.SetConnMaxLifetime(time.Minute * 5) // 5分钟后关闭空闲连接

	//数据库迁移操作（升级或者字段变化）
	//m, err := migrate.New("file://path/to/migrations", "sqlite://path/to/database.db")
	//if err != nil {
	//	log.Fatalf("Failed to initialize migrations: %v", err)
	//}
	//
	//err = m.Up()
	//if err != nil && err != migrate.ErrNoChange {
	//	log.Fatalf("Failed to apply migrations: %v", err)
	//}

	//设置缓存大小
	cacheSize := 100 * 1024 * 1024 // 100MB缓存大小
	ds.cache = freecache.NewCache(cacheSize)

	return err
}

func (ds *DatabaseStruct) CloseSqlite() error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if ds.cache != nil {
		ds.cache.Clear()
	}

	if ds.sqlDB != nil {
		if err := ds.sqlDB.Close(); err == nil {
			ds.db = nil
		} else {
			return err
		}
	}
	return nil
}

func (ds *DatabaseStruct) Create(tableName string, v interface{}, structEntity any) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if ds.db != nil {
		// 运行自动迁移创建表结构（假设structEntity模型对应的表不存在，GORM会尝试创建）
		ds.db.AutoMigrate(structEntity)
		return ds.db.Table(tableName).Create(v).Error
	}
	return nil
}

func (ds *DatabaseStruct) Delete(tableName string, ids []string, structEntity any) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if ds.db != nil && ids != nil {
		return ds.db.Table(tableName).Where("id IN ?", ids).Delete(structEntity).Error
	}
	return nil
}

func (ds *DatabaseStruct) Find(tableName string, ids []string, structEntity any) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if ds.db != nil {
		if ids != nil && len(ids) != 0 {
			//// 尝试从缓存中获取
			//userBytes, err := cache.Get([]byte(fmt.Sprintf("user:%d", id)))
			//if err == nil {
			//	// 缓存命中
			//	var user User
			//	json.Unmarshal(userBytes, &user)
			//	return &user, nil
			//}

			return ds.db.Table(tableName).Where("id IN ?", ids).Find(structEntity).Error
		} else {
			return ds.db.Table(tableName).Find(structEntity).Error
		}
	}
	return nil
}

func (ds *DatabaseStruct) FindColumns(tableName string, structEntity any, query interface{}, args ...interface{}) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if ds.db != nil {
		if query != nil && len(args) != 0 {
			//// 尝试从缓存中获取
			//userBytes, err := cache.Get([]byte(fmt.Sprintf("user:%d", id)))
			//if err == nil {
			//	// 缓存命中
			//	var user User
			//	json.Unmarshal(userBytes, &user)
			//	return &user, nil
			//}

			return ds.db.Table(tableName).Where(query, args...).Find(structEntity).Error
		} else {
			return ds.db.Table(tableName).Find(structEntity).Error
		}
	}
	return nil
}

func (ds *DatabaseStruct) Update(tableName string, id uint, structEntity any) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if ds.db != nil && id != 0 {
		return ds.db.Table(tableName).Updates(structEntity).Error
	}
	return nil
}
