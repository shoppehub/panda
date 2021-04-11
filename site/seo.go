package site

import (
	"gorm.io/gorm"
)

// SEO 配置
type Seo struct {
	gorm.Model
	SiteId string `json:"siteId,omitempty"`
	// 页面的 ID
	PageId string `json:"pageId,omitempty"`
	// 页面的 title
	Title string `json:"title,omitempty"`
	// 页面的 描述
	Desc string `json:"desc,omitempty"`
	// 页面的 关键字
	Keyword string `json:"keyword,omitempty"`
	// url 别名
	UrlAlias string `json:"urlAlias,omitempty"`
}

// SEO 配置
type SeoRedirect struct {
	gorm.Model
	SiteId string `json:"siteId,omitempty"`
	// 页面的 ID
	SourcePath string `json:"sourcePath,omitempty"`
	// 页面的 title
	Target string `json:"target,omitempty"`
	// 跳转码，301或者302
	Status int `json:"status,omitempty"`
}
