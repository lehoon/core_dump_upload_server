@echo off
@rem echo ##!/usr/bin/env sh
@rem echo set CGO_ENABLE=1
@rem echo set GOOS=linux
@rem echo set GOARCH=amd64

@rem echo go build -ldflags "-s -w"
go build -ldflags "-w"

@rem echo "给go程序增加图标"
@rem echo "step1 首先在cmd下使用 windres -o hook_api.syso hook_api.rc   生成syso文件"
@rem echo "step2 再使用go build编译就会生成带图标的exe或者elf程序"

