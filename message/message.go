package message

// 设备信息
type AppDumpInfo struct {
	AppId      string `json:"appId"`
	Version    string `json:"version"`
	FilePath   string `json:"filepath"`
	RemoteHost string `json:"remotehost"`
	Created    string `json:"created"`
}

func (d *AppDumpInfo) IsEmpty() bool {
	return len(d.AppId) == 0 && len(d.Version) == 0 && len(d.FilePath) == 0 && len(d.RemoteHost) == 0
}
