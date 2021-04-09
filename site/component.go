package site

// 组件定义
type Component struct {
	// 组件id
	ComponentId string `json:"componentId,omitempty"`
	/**
	匹配多个区域类型，如果是 * 则全匹配
	*/
	MatchRegionType []string `json:"type,omitempty"`
	// 自定义样式
	StyleData map[string]interface{} `json:"styleData,omitempty"`
	// 自定义数据
	Data map[string]interface{} `json:"data,omitempty"`
}
