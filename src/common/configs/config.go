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
	WSCostModels   WSCostModels   `json:"WS_Cost_Models"`
	WSFiles        WSFiles        `json:"WS_Files"`
	WSTasks        WSTasks        `json:"WS_Tasks"`
}

type WSProjects struct {
	WSCreateForTask string `json:"Create_For_Task"`
	WSCancelProject string `json:"Cancel_Project"`
}

type WSProjectGroup struct {
	ProjectGroupCreate string `json:"ProjectGroup_Create_API"`
}
type WSInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WSCostModels struct {
	CostModels string `json:"Cost_Models"`
}

type WSFiles struct {
	Files string `json:"Files"`
}

type WSTasks struct {
	Tasks string `json:"Tasks"`
}
