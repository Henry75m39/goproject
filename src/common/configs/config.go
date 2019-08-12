package configs

const (
	ConfigPrefix = "class/"
	Namespace    = "com.vmware."
	LogPath      = "./data/goapp/log/"
)

const (
	ServiceNameTMS = "tms"
)

type WSConfig struct {
	WSLoginAPI                  string `json:"WS_Login_API"`
	WSTMSLookupAPI              string `json:"WS_TMSLookup_API"`
	WSTMSProjectGroupsCreateAPI string `json:"WS_TMS_ProjectGroups_Create_API"`

	WSInfo WSInfo `json:"WS_Info"`
}

type WSInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
