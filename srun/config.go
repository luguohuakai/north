package srun

// Northbound @description SDK配置 使用到项目中时酌情修改
// @author DM
// @time 2022/04/17 00:11
type Northbound struct {
	Protocol    string `json:"Protocol,omitempty"`
	InterfaceIp string `json:"InterfaceIp,omitempty"`
	Port        int    `json:"Port,omitempty"`
}
