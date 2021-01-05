package domain

import (
	"context"
	"fmt"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reflect"
)

const TransactionContext = "DB_CTX"

type Database struct {
	DSN     string
	Session *gorm.DB
	Config  *gorm.Config
}

var DB Database

// 初始化数据库
// 建立数据库连接、ORM，并执行外部传入的脚本进行
func InitDB(scripts ...string) *gorm.DB {
	DB.DSN = infrasturcture.GetConfiguration().DSN
	DB.Config = &gorm.Config{}
	db, err := gorm.Open(sqlite.Open(DB.DSN), DB.Config)
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}
	// 处理初始化DDL、DML脚本
	for _, v := range scripts {
		if len(v) > 0 {
			db.Exec(v)
		}
	}
	DB.Session = db
	return db
}

// 取当前访问上下文中的数据库连接
// 当前上下文中，由事务中间件自动开启和提交事务，无需手动管理
func (db *Database) Context(ctx context.Context) *Database {
	session := ctx.Value(TransactionContext).(*gorm.DB)
	ctxDB := new(Database)
	*ctxDB = *db
	ctxDB.Session = session
	return ctxDB
}

func (db *Database) String() string {
	return fmt.Sprintf("Conn: %v@%v", reflect.TypeOf(db.Session.Statement.ConnPool), &db.Session.Statement.ConnPool)
}
