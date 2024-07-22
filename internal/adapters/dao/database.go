// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package dao

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/common"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

type IDatabase interface {
	// 定义所有数据库操作的公共方法签名
	// ...
}

type DatabaseStruct struct {
	db     *gorm.DB
	mu     sync.Mutex
	logger *zap.Logger

	Gen *Query
}

func OpenDB(cfg config.AppConfig) (*DatabaseStruct, error) {
	var err error = nil
	var dialector gorm.Dialector = nil

	// 创建一个默认的日志器实例，并设置输出级别为 logger.Info
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	switch cfg.Database.DBtype {
	case "postgresql":
		dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Shanghai",
			cfg.Database.DBtype, cfg.Database.User, cfg.Database.Password, cfg.Database.Host,
			cfg.Database.Port, cfg.Database.DBName)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dsn := "./db/sqlite/sqlite.db" //我已经导入初始数据到sqlite文件里，放在项目上了
		dialector = sqlite.Open(dsn)
	default: //mysql
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.User,
			cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
		dialector = mysql.Open(dsn)
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})
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

	// 初始化自动生成的组件
	g := Use(db)

	return &DatabaseStruct{
		db:  db,
		Gen: g,
	}, nil
}

func (ds *DatabaseStruct) CloseDB() error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

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
	ds.mu.Lock()
	defer ds.mu.Unlock()

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

// 泛型居然在我的电脑不支持
//func (ds *DatabaseStruct) paginate(pageReq common.PageRequest) (*common.PageResponse[interface{}], error) {
//	var total int64
//	ds.db.Model(&T{}).Count(&total)
//
//	if pageReq.PageSize == 0 {
//		pageReq.PageSize = 10 // 默认每页显示10条数据
//	}
//
//	if pageReq.PageNum < 1 {
//		pageReq.PageNum = 1 // 默认第一页
//	}
//
//	var data []interface{}
//	ds.db.Offset((pageReq.PageNum - 1) * pageReq.PageSize).Limit(pageReq.PageSize).Find(&data)
//
//	return &common.PageResponse[T]{
//		Total:    int(total),
//		PageNum:  pageReq.PageNum,
//		PageSize: pageReq.PageSize,
//		Data:     data,
//	}, nil
//}

func (ds *DatabaseStruct) PageQuery(txFunc func(*gorm.DB) *gorm.DB, pageReq common.PageRequest, results any) (int64, error) {
	var total int64
	err := txFunc(ds.db).Count(&total).Error
	if err != nil {
		return 0, err
	}

	err = txFunc(ds.db).Limit(pageReq.PageSize).Offset((pageReq.PageNum - 1) * pageReq.PageSize).Scan(results).Error
	if err != nil {
		return total, err
	}

	return total, nil
}
