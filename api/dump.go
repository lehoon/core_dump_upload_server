package api

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/render"
	"github.com/lehoon/core_dump_upload_server/v2/library/logger"
	los "github.com/lehoon/core_dump_upload_server/v2/library/os"
	"github.com/lehoon/core_dump_upload_server/v2/message"
	"github.com/lehoon/core_dump_upload_server/v2/service"
)

func UploadDumpFile(w http.ResponseWriter, r *http.Request) {
	logger.Log().Info("接收到新的程序dump上传请求.")

	//检查流是否已经配置
	appId := r.URL.Query().Get("appId")
	if appId == "" {
		logger.Log().Info("未携带有效的appId参数.")
		render.Respond(w, r, FailureBizResultWithParamError())
		return
	}

	version := r.URL.Query().Get("version")
	if version == "" {
		logger.Log().Info("未携带有效的version参数.")
		render.Respond(w, r, FailureBizResultWithParamError())
		return
	}

	//是否启用gzip压缩
	contentEncoding := r.Header.Get("Content-Encoding")
	if contentEncoding == "gzip" {
		r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 设置最大请求体大小为 10MB
		reader, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, "Failed to create gzip reader.", http.StatusInternalServerError)
			return
		}
		defer reader.Close()

		r.Body = http.MaxBytesReader(w, reader, 10<<20) // 设置最大请求体大小为 10MB
	}

	// 解析表单，设置边界
	err := r.ParseMultipartForm(32 << 20) // 32MB 边界大小
	if err != nil {
		logger.Log().Info("保存上传的dump文件失败.")
		http.Error(w, "Failed to parse multipart form.", http.StatusBadRequest)
		return
	}

	// 从请求中读取文件
	file, handler, err := r.FormFile("upload_file_minidump")
	if err != nil {
		logger.Log().Info("读取上传的dump文件信息失败.")
		http.Error(w, "Failed to read file from request.", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf(handler.Filename)
	// 创建一个本地文件来保存上传的文件
	if !los.IsFileExist("./uploads/" + appId) {
		los.Mkdir("./uploads/" + appId)
	}

	uploadedFile, err := os.Create(filepath.Join("./uploads/"+appId, handler.Filename))
	if err != nil {
		logger.Log().Info("创建dump文件失败{}", handler.Filename)
		http.Error(w, "Failed to create file on the server.", http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	// 将上传的文件内容拷贝到本地文件
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		logger.Log().Infof("拷贝dump文件内容到新文件失败%s", handler.Filename)
		http.Error(w, "Failed to save file on the server.", http.StatusInternalServerError)
		return
	}

	//创建数据库记录
	record := &message.AppDumpInfo{
		AppId:      appId,
		Version:    version,
		FilePath:   uploadedFile.Name(),
		RemoteHost: r.RemoteAddr,
	}
	service.InsertDumpInfo(record)
	logger.Log().Infof("创建dump文件成功%s", handler.Filename)
	render.Respond(w, r, SuccessBizResult())
}
