package model

import (
	"time"
)

// 默认的时间格式
const JsonTimeLayout = "2006-01-02 15:04:05"

// 时间模型，设置开始和结束时间
type Gmt struct {
	GmtCreated  time.Time `json:"gmtCreated,omitempty"`
	GmtModified time.Time `json:"gmtModified,omitempty"`
}

// id 为 string 的基础模型
type BaseStringId struct {
	Id string `json:"id,omitempty"`
	Gmt
}
