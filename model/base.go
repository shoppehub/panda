package model

import (
	"time"
)

type JsonTime time.Time

const JsonTimeLayout = "2006-01-02 15:04:05"

type gmt struct {
	GmtCreated  *JsonTime `json:"gmtCreated,omitempty"`
	GmtModified *JsonTime `json:"gmtModified,omitempty"`
}

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var tt = time.Time(this)

	if tt.IsZero() {
		return []byte("null"), nil
	}

	data := make([]byte, 0)
	data = append(data, '"')
	data = tt.AppendFormat(data, JsonTimeLayout)
	data = append(data, '"')

	return data, nil
}

// 实现它的json反序列化方法
func (this *JsonTime) UnmarshalJSON(data []byte) error {

	if data == nil {
		return nil
	}

	ts, err := time.Parse(`"`+JsonTimeLayout+`"`, string(data))

	//this = time
	*this = JsonTime(ts)

	return err
}

// 时间模型，设置开始和结束时间
type Gmt struct {
	GmtCreated  JsonTime `json:"gmtCreated,omitempty"`
	GmtModified JsonTime `json:"gmtModified,omitempty"`
}

// id 为 string 的基础模型
type BaseStringId struct {
	Id string `json:"id,omitempty"`
	Gmt
}
