package os

import (
	"os"
	"path/filepath"
)

// IsJavaInstall 检查java环境  true  已经安装  false 没有安装
func IsJavaInstall() bool {
	//查询env变量是否有JAVA_HOME
	javaHome := os.Getenv("JAVA_HOME")
	if len(javaHome) == 0 {
		return false
	}

	javaBinPath := javaHome + string(filepath.Separator) + "bin" + string(filepath.Separator) + "java.exe"
	if IsFileExist(javaBinPath) {
		return true
	}

	//根据JAVA_HOME找java.exe
	return false
}
