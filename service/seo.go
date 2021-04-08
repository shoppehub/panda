package service

import "github.com/shoppehub/shoppe/model"

// SEO 配置
type Seo struct {
	model.BaseStringId
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
	model.BaseStringId
	SiteId string `json:"siteId,omitempty"`
	// 页面的 ID
	SourcePath string `json:"sourcePath,omitempty"`
	// 页面的 title
	Target string `json:"target,omitempty"`
	// 跳转码，301或者302
	Status int `json:"status,omitempty"`
}
