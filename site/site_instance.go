package site

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
	// // 导航信息
	// Navs []SiteNav `json:"navs,omitempty"`

	// 页面实例
	Pages []SitePage `json:"pages,omitempty"`
	// 模板id
	TemplateId string `json:"templateId,omitempty"`
}

// // 站点导航
// type SiteNav struct {
// 	model.BaseStringId
// 	// 站点id
// 	SiteId string `json:"siteId,omitempty"`
// 	// 站点实例Id
// 	SiteInstanceId string `json:"siteInstanceId,omitempty"`
// 	// 页面名称
// 	PageTitle string `json:"pageName,omitempty"`
// 	// 页面URL
// 	PageUrl string `json:"pageUrl,omitempty"`
// 	// 关联的页面ID
// 	PageInstanceId string `json:"pageInstanceId,omitempty"`
// 	// 排序
// 	Sort int `json:"sort,omitempty"`
// }
