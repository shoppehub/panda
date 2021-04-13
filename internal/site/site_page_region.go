package site

// 站点段落内区域实例
type SitePageRegion struct {
	// 列的序号，用于控制位置
	ColumnIndex int `json:"columnIndex,omitempty"`
	// 自定义样式
	StyleData map[string]interface{} `json:"styleData,omitempty"`
	//区块内的组件
	Components []SitePageComponent `json:"components,omitempty"`
}
