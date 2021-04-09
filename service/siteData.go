package service

import "github.com/shoppehub/shoppe/model"

//站点实例
type SiteInstance struct {
	model.BaseStringId
	// 站点id
	SiteId string `json:"siteId,omitempty"`
	// 实例版本，保存一次就增加一次version
	Version int `json:"version,omitempty"`
	// 版本别名，用于支持用户自定义名称
	Alias string `json:"alias,omitempty"`
	// 导航信息
	Navs []SiteNav `json:"navs,omitempty"`
	// 页面实例
	Pages []SitePage `json:"pages,omitempty"`
}

// 站点导航
type SiteNav struct {
	model.BaseStringId
	// 站点id
	SiteId string `json:"siteId,omitempty"`
	// 站点实例Id
	SiteInstanceId string `json:"siteInstanceId,omitempty"`
	// 页面名称
	PageTitle string `json:"pageName,omitempty"`
	// 页面URL
	PageUrl string `json:"pageUrl,omitempty"`
	// 关联的页面ID
	PageInstanceId string `json:"pageInstanceId,omitempty"`
	// 排序
	Sort int `json:"sort,omitempty"`
}

// 站点页面实例
type SitePage struct {
	model.BaseStringId
	SiteInstanceId string `json:"siteInstanceId,omitempty"`
	// 页面名称，一般用于系统页面
	PageName string `json:"pageName,omitempty"`
	// 页面标题
	PageTitle string `json:"pageTitle,omitempty"`
	// 页面主题
	Theme string `json:"theme,omitempty"`
	// 页面是否开启，默认开启，不开启相当于禁用页面，自动跳转会网站首页
	Enable bool `json:"enable,omitempty"`
	// 段落配置
	Segments []SitePageSegment `json:"segments,omitempty"`
}

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
	//区块，每个
	Regions []SitePageSegment `json:"regions,omitempty"`
}

type SitePageComponent struct {
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
