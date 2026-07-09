package model

type MonitorServer struct {
	Cpu      *CpuInfo   `json:"cpu"`
	Mem      *MemInfo   `json:"mem"`
	Jvm      *JvmInfo   `json:"jvm"`
	Sys      *SysInfo   `json:"sys"`
	SysFiles []*SysFile `json:"sysFiles"`
}

type CpuInfo struct {
	CpuNum int     `json:"cpuNum"`
	Total  float64 `json:"total"`
	Sys    float64 `json:"sys"`
	Used   float64 `json:"used"`
	Wait   float64 `json:"wait"`
	Free   float64 `json:"free"`
}

type MemInfo struct {
	Total float64 `json:"total"`
	Used  float64 `json:"used"`
	Free  float64 `json:"free"`
	Usage float64 `json:"usage"`
}

type JvmInfo struct {
	Total     float64 `json:"total"`
	Max       float64 `json:"max"`
	Free      float64 `json:"free"`
	Used      float64 `json:"used"`
	Usage     float64 `json:"usage"`
	Name      string  `json:"name"`
	Version   string  `json:"version"`
	Home      string  `json:"home"`
	StartTime string  `json:"startTime"`
	RunTime   string  `json:"runTime"`
	InputArgs string  `json:"inputArgs"`
}

type SysInfo struct {
	ComputerName string `json:"computerName"`
	ComputerIp   string `json:"computerIp"`
	UserDir      string `json:"userDir"`
	OsName       string `json:"osName"`
	OsArch       string `json:"osArch"`
}

type SysFile struct {
	DirName     string  `json:"dirName"`
	SysTypeName string  `json:"sysTypeName"`
	TypeName    string  `json:"typeName"`
	Total       string  `json:"total"`
	Free        string  `json:"free"`
	Used        string  `json:"used"`
	Usage       float64 `json:"usage"`
}
