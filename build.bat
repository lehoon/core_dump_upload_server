@echo off
@rem echo ##!/usr/bin/env sh
@rem echo set CGO_ENABLE=1
@rem echo set GOOS=linux
@rem echo set GOARCH=amd64

@rem echo go build -ldflags "-s -w"
go build -ldflags "-w"

@rem echo "��go��������ͼ��"
@rem echo "step1 ������cmd��ʹ�� windres -o hook_api.syso hook_api.rc   ����syso�ļ�"
@rem echo "step2 ��ʹ��go build����ͻ����ɴ�ͼ���exe����elf����"

