package site

import "time"

// 站点模型,前台和后台公用一个模型
type Site struct {
	ID        uint `gorm:"primarykey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// 二级域名
	Subdomain string `json:"subdomain,omitempty"`
	// 顶级域名
	Domain string `json:"domain,omitempty"`
	// 用户Id
	UserId uint `json:"userId,omitempty"`
	// 站点名称
	Name string `json:"name,omitempty"`
	// 站点分类
	CategoryId string `json:"categoryId,omitempty"`
	// 站点版本
	Version int `json:"version,omitempty"`
	// 网站 favicon
	Favicon string `json:"favicon,omitempty"`
}

// 初始化站点
func (s *Site) InitSite() *Site {

	return nil
}
