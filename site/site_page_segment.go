package site

// 站点段落实例
type SitePageSegment struct {
	/**
	段落类型，用于设置段落的布局,
	full-column 全屏
	center-column 通栏，比如说浏览器宽度是 1440，通栏就是1200像素并且居中的效果
	left-main-column 左右布局
	left-main-right-column 左中右布局
	main-right-column 中右布局
	*/
	Layout string `json:"layout,omitempty"`
	// 自定义样式
	StyleData map[string]interface{} `json:"styleData,omitempty"`
	//区块，每个
	Regions []SitePageSegment `json:"regions,omitempty"`
}

// 站点段落内区域实例
type SitePageRegion struct {
	// 列的序号，用于控制位置
	ColumnIndex int `json:"columnIndex,omitempty"`
	// 自定义样式
	StyleData map[string]interface{} `json:"styleData,omitempty"`
	//区块内的组件
	Components []Component `json:"components,omitempty"`
}
