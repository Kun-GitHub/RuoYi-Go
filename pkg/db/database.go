// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package rydb

import (
	"RuoYi-Go/config"
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
	Cache *freecache.Cache

	mu sync.Mutex
}

func OpenSqlite(cfg config.AppConfig) (*DatabaseStruct, error) {
	var err error = nil
	var dialector gorm.Dialector = nil

	switch cfg.Database.DBtype {
	case "postgresql":
		dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Shanghai",
			cfg.Database.DBtype, cfg.Database.User, cfg.Database.Password, cfg.Database.Host,
			cfg.Database.Port, cfg.Database.DBName)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dsn := "./test.db"
		dialector = sqlite.Open(dsn)
	default:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.User,
			cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
		dialector = mysql.Open(dsn)
	}
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数（例如设置最大空闲连接数和最大打开连接数）
	sqlDB.SetMaxIdleConns(10)  // 设置空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100) // 设置数据库的最大打开连接数

	// 如果需要，还可以设置连接在空闲多久后关闭（单位秒）
	sqlDB.SetConnMaxLifetime(time.Minute * 5) // 5分钟后关闭空闲连接

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
	cache := freecache.NewCache(cacheSize)

	return &DatabaseStruct{
		db:    db,
		Cache: cache,
	}, nil
}

func (ds *DatabaseStruct) CloseSqlite() error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if ds.Cache != nil {
		ds.Cache.Clear()
	}

	if ds.db != nil {
		sqlDB, err := ds.db.DB()
		if err == nil {
			if err := sqlDB.Close(); err == nil {
				ds.db = nil
			} else {
				return err
			}
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
			return ds.db.Table(tableName).Where(query, args...).Find(structEntity).Error
		} else {
			return fmt.Errorf("没有传查询参数")
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

func (ds *DatabaseStruct) Transactional(txFunc func(*gorm.DB) error) error {
	tx := ds.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		if err := recover(); tx.Error != nil || err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	if err := txFunc(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (ds *DatabaseStruct) CustomQuery(query string, args []interface{}, processRow func(rows *sql.Rows) error) error {
	var rows *sql.Rows
	var err error
	if rows, err = ds.db.Raw(query, args...).Rows(); err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := processRow(rows); err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
