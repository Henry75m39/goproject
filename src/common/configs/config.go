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
	WSLoginAPI     string         `json:"WS_Login_API"`
	WSTMSLookupAPI string         `json:"WS_TMSLookup_API"`
	WSProjects     WSProjects     `json:"WS_Projects"`
	WSProjectGroup WSProjectGroup `json:"WS_Project_Group"`
	WSInfo         WSInfo         `json:"WS_Info"`
}

type WSProjects struct {
	WSCreateForTask string `json:"Create_For_Task"`
}

type WSProjectGroup struct {
	ProjectGroupCreate string `json:"ProjectGroup_Create_API"`
}
type WSInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
