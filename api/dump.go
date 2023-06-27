package api

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/render"
	"github.com/lehoon/core_dump_upload_server/v2/library/logger"
)

func UploadDumpFile(w http.ResponseWriter, r *http.Request) {
	logger.Log().Info("stream no reader event received")

	//检查流是否已经配置
	appId := r.URL.Query().Get("id")
	if appId == "" {
		render.Respond(w, r, FailureBizResultWithParamError())
		return
	}

	version := r.URL.Query().Get("version")
	if version == "" {
		render.Respond(w, r, FailureBizResultWithParamError())
		return
	}

	guid := r.URL.Query().Get("guid")
	if guid == "" {
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

	// 创建一个本地文件来保存上传的文件
	uploadedFile, err := os.Create(filepath.Join("./uploads", guid+".dmp"))
	if err != nil {
		http.Error(w, "Failed to create file on the server.", http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	// 将上传的文件内容拷贝到本地文件
	_, err = io.Copy(uploadedFile, r.Body)
	if err != nil {
		http.Error(w, "Failed to save file on the server.", http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, SuccessBizResult())
}
