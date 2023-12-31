package database

import (
	"database/sql"

	"github.com/lehoon/core_dump_upload_server/v2/library/logger"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbOpenState         = false
	dbInstance  *sql.DB = nil
)

var (
	NoResultError = sql.ErrNoRows
)

func init() {
	var err error = nil
	dbInstance, err = sql.Open("sqlite3", "storage.db")
	if err != nil {
		logger.Log().Errorf("打开数据库文件失败, %v", err.Error())
	}

	dbOpenState = true
	logger.Log().Info("打开数据库storage.db成功")
}

func IsOpen() bool {
	return dbOpenState
}

func Instance() *sql.DB {
	return dbInstance
}

func Shutdown() {
	if !dbOpenState {
		return
	}

	dbInstance.Close()
}
