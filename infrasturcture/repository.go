package infrasturcture

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reflect"
)

type Database struct {
	DSN     string
	Session *gorm.DB
	Config  *gorm.Config
}

var DB Database

// 初始化数据库
// 建立数据库连接、ORM，并执行外部传入的脚本进行
func InitDB(scripts ...string) *gorm.DB {
	DB.DSN = GetConfiguration().DSN
	DB.Config = &gorm.Config{
		Logger: GORMLogger(),
	}
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

func TxDB(ctx context.Context) *Database {
	ctxDB := new(Database)
	*ctxDB = DB
	// 如果上下文中有数据库会话，就使用上下文的，否则就使用全局的
	if session, ok := ctx.Value(CTXTransaction).(*gorm.DB); ok {
		ctxDB.Session = session
	}
	return ctxDB
}

func TxID(st *gorm.Statement) string {
	return fmt.Sprintf("%v@%v", reflect.TypeOf(st.ConnPool), &st.ConnPool)
}

func (db *Database) String() string {
	return TxID(db.Session.Statement)
}
