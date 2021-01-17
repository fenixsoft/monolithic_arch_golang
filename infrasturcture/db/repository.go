package db

import (
	"fmt"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/config"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reflect"
)

const CTXDatabase = "CTX_DB"

type TransactionState string

const (
	TransactionStateGlobal     = "Non-Transaction" // 全局连接，未开启事务
	TransactionStateProcessing = "Processing"      // 已开启事务，正在处理中
	TransactionStateCommit     = "Commit"          // 事务已提交
	TransactionStateRollback   = "Rollback"        // 事务已回滚
)

type Database struct {
	DSN     string
	Session *gorm.DB
	Config  *gorm.Config
	State   TransactionState
	Error   error
}

type Result struct {
	Error        error
	RowsAffected int64
}

var DB Database

// 初始化数据库
// 建立数据库连接、ORM，并执行外部传入的脚本进行
func InitDB(scripts ...string) *gorm.DB {
	DB.DSN = config.GetConfiguration().DSN
	DB.Config = &gorm.Config{
		Logger:                 logger.GORMLogger(),
		SkipDefaultTransaction: true,
	}
	DB.State = TransactionStateGlobal
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

func NewDB() *Database {
	newDB := new(Database)
	*newDB = DB
	return newDB
}

func NewWithTx(tx *gorm.DB) *Database {
	txDB := new(Database)
	*txDB = DB
	txDB.Session = tx
	txDB.State = TransactionStateProcessing
	return txDB
}

func TxID(st *gorm.Statement) string {
	return fmt.Sprintf("%v@%v", reflect.TypeOf(st.ConnPool), &st.ConnPool)
}

func (db *Database) String() string {
	return TxID(db.Session.Statement)
}
