package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/lehoon/core_dump_upload_server/v2/library/logger"
)

func RequestLoggerFilter(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		buff, _ := io.ReadAll(r.Body) // , string(buff)
		fmt.Printf("Request url=%s,method=%s,body=\n", r.RequestURI, r.Method)
		logger.Log().Infof("Request url=%s,method=%s,body=\n", r.RequestURI, r.Method)

		//把读出来的数据再写到request.body上
		r.Body = io.NopCloser(bytes.NewBuffer(buff))
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
