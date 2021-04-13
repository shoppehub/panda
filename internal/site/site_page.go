package site

import "gorm.io/gorm"

// 站点页面实例
type SitePage struct {
	gorm.Model
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
