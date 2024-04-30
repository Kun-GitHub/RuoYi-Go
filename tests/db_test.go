package tests

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time-machine/internal/db"
)

// 定义模型（示例为User模型）
type User struct {
	gorm.Model
	Name  string
	Email string
}

func TestSqlite(t *testing.T) {

	// 创建sqliteStruct实例
	sqliteService := &db.DatabaseStruct{
		DatabaseType: "sqlite",
		Dsn:          "./test.db", // SQLite 数据库文件路径
	}

	// 打开数据库并执行自动迁移（假设User结构体是你的模型）
	err := sqliteService.OpenSqlite()
	if err != nil {
		fmt.Println("Failed to open or migrate the database:", err)
		return
	}

	// 在这里执行数据库相关操作，如查询、插入、更新等
	// ...
	//sqliteService.Create(&User{Name: "张三", Email: "zhangsan@example.com"})

	// 完成所有操作后，关闭数据库连接
	defer func() {
		err = sqliteService.CloseSqlite()
		if err != nil {
			fmt.Println("Failed to close the database connection:", err)
		}
	}()

}
