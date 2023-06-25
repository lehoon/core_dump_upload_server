package api

import (
	"net/http"

	"github.com/lehoon/core_dump_upload_server/v2/library/logger"
)

func QueryOperateCodeMessage(w http.ResponseWriter, r *http.Request) {
	//data := OperateCodeMessage()
	//render.Respond(w, r, SuccessBizResultWithData(data))
}

func ShowPostMessage(w http.ResponseWriter, r *http.Request) {
	logger.Log().Infof("ShowPostMessage   %s", r.Body)
	//render.Respond(w, r, SuccessBizResultWithData(data))
}
