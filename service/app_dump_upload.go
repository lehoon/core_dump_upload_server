package service

import (
	"errors"
	"strings"

	"github.com/lehoon/core_dump_upload_server/v2/library/database"
	"github.com/lehoon/core_dump_upload_server/v2/library/logger"
	"github.com/lehoon/core_dump_upload_server/v2/library/utils"
	"github.com/lehoon/core_dump_upload_server/v2/message"
)

func init() {
	logger.Log().Info("app dump upload service initialized")
}

// 新增app dump info
func InsertDumpInfo(record *message.AppDumpInfo) error {
	if !database.IsOpen() {
		return errors.New("数据库未打开,新增失败")
	}

	streamid_max, err := sequence_no_max()
	if err != nil {
		logger.Log().Errorf("新增app生成的dump文件记录失败,数据库发送错误, %s,%s", utils.JsonString(record), err.Error())
		return errors.New("新增app生成的dump文件记录失败,请稍后重试")
	}

	streamid, err := next_sequece()
	if err != nil {
		logger.Log().Errorf("新增app生成的dump文件记录失败,获取流序列号失败, %s,%s", utils.JsonString(record), err.Error())
		return errors.New("新增app生成的dump文件记录失败,请稍后重试")
	}

	streamid_new_int := string_to_int32(streamid)

	if streamid_new_int <= streamid_max {
		streamid_new_int = streamid_max + 1
		_, err = update_sequence(streamid_new_int)
		if err != nil {
			logger.Log().Errorf("新增app生成的dump文件记录失败,获取流序列号失败, %s,%s", utils.JsonString(record), err.Error())
			return errors.New("新增app生成的dump文件记录失败,请稍后重试")
		}
	}

	//添加新的数据
	insertSql := `insert into app_dump_upload(id,appid,version,filepath,remotehost,created) values(?,?,?,?,?,datetime('now','localtime'))`
	insertStmt, err := database.Instance().Prepare(insertSql)

	if err != nil {
		logger.Log().Errorf("新增app生成的dump文件记录失败,%s,%s", utils.JsonString(record), err.Error())
		return errors.New("新增app生成的dump文件记录失败,请稍后重试")
	}

	defer insertStmt.Close()
	_, err = insertStmt.Exec(streamid_new_int, record.AppId, record.Version, record.FilePath, record.RemoteHost)

	if err != nil {
		logger.Log().Errorf("新增app生成的dump文件记录数据失败,%s,%s", utils.JsonString(record), err.Error())
		return errors.New("新增app生成的dump文件记录失败,请稍后重试")
	}

	return nil
}

// 查询流最大编号
func sequence_no_max() (int32, error) {
	if !database.IsOpen() {
		return 0, errors.New("查询app生成的dump文件记录表流序号失败,当前数据库未建立连接")
	}

	if count_app_dump_table() == 0 {
		return 0, nil
	}

	stmt, err := database.Instance().Prepare("select max(cast(id as int)) as id from app_dump_upload")
	if err != nil {
		logger.Log().Error("查询app生成的dump文件记录表流序号失败, %s", err.Error())
		return 0, errors.New("查询app生成的dump文件记录表流序号失败,请稍后重试")
	}

	defer stmt.Close()
	row := stmt.QueryRow()

	var streamid string
	err = row.Scan(&streamid)

	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		logger.Log().Errorf("生成流序列号失败,未初始化数据,%s", err.Error())
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return string_to_int32(streamid), nil
}

// 查询app生成的dump文件记录表数量
func count_app_dump_table() int {
	if !database.IsOpen() {
		return 0
	}

	stmt, err := database.Instance().Prepare("select count(*) as totalcount from app_dump_upload")
	if err != nil {
		logger.Log().Errorf("检查app生成的dump文件记录表失败,%s", err.Error())
		return 0
	}

	defer stmt.Close()
	row := stmt.QueryRow()

	var totalcount int
	err = row.Scan(&totalcount)

	if err != nil {
		return 0
	}

	return totalcount
}
