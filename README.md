# dump文件服务器
crashpad上传dump文件，文件服务器，提供收集dump文件、触发报警功能，文件下载功能.

## 收集dump文件功能
* 对外提供dump文件上传功能，/api/v1/dump/upload?appId=&version=, 客户端配置好url，即可在生成dump文件的时候把dump文件上传到该服务器。
文件上传目前适配crashpad上传逻辑，不启用gzip压缩功能。

* 文件按照appid/version 目录保存dump文件。

## 报警功能
当接收到客户端上传的dump文件后，触发报警功能，通过http协议把一个报警信息发送到业务平台。
报警信息包括：appId信息，版本信息，时间戳，文件下载路径等信息。

## 文件下载功能
开发人员通过报警信息，拿到下载地址后，下载dump文件，与指定版本号的pdb文件调试，接近引起dump的代码问题。





