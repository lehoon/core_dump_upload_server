package service

import (
	"github.com/lehoon/core_dump_upload_server/v2/library/database"
	"github.com/lehoon/core_dump_upload_server/v2/library/logger"
)

func init() {
	createDumpTable()
	createSequenceTable()
	logger.Log().Info("database service initialized")
}

// 创建app_dump表
func createDumpTable() {
	if !database.IsOpen() {
		logger.Log().Info("创建设备表失败,当前数据库未建立连接")
		panic("创建设备表失败,当前数据库未建立连接")
	}

	appDumpUploadTableSql := `create table if not exists app_dump_upload (
		 id         INTEGER PRIMARY KEY,
         appid      varchar(32) NOT NULL,
		 version    varchar(32) NOT NULL,
         filepath   varchar(32) NOT NULL,
		 remotehost varchar(32) NOT NULL,
         created    DATE NOT NULL
         );
    `

	_, err := database.Instance().Exec(appDumpUploadTableSql)
	if err != nil {
		logger.Log().Errorf("创建表结构app_dump_upload失败 %v", err)
		return
	}
}

// 创建自增序列表维护一个流序号
func createSequenceTable() {
	if !database.IsOpen() {
		logger.Log().Info("创建自增序列表失败,当前数据库未建立连接")
		panic("创建自增序列表失败,当前数据库未建立连接")
	}

	sequenceTableSql := `create table if not exists sequence_info (
         sequenceid integer );
    `

	_, err := database.Instance().Exec(sequenceTableSql)
	if err != nil {
		logger.Log().Errorf("创建表结构sequence_info失败 %v", err)
		return
	}
}
