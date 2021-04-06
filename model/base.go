package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonTime time.Time

const JsonTimeLayout = "2006-01-02 15:04:05"

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var tt = time.Time(this)

	if tt.IsZero() {
		fmt.Println(tt)
		return []byte("null"), nil
	}

	var stamp = fmt.Sprintf("\"%s\"", tt.Format(JsonTimeLayout))

	return []byte(stamp), nil
}

// 实现它的json反序列化方法
func (this *JsonTime) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	ts, err := time.Parse(JsonTimeLayout, string(data))
	//this = time
	*this = JsonTime(ts)
	return err
}

// 实现它的json序列化方法
func (this Gmt) MarshalJSON() ([]byte, error) {

	created := JsonTime(this.GmtCreated)
	modified := JsonTime(this.GmtModified)

	return json.Marshal(struct {
		GmtCreated  *JsonTime `json:"gmtCreated,omitempty"`
		GmtModified *JsonTime `json:"gmtModified,omitempty"`
	}{
		GmtCreated:  &created,
		GmtModified: &modified,
	})
}

// 实现它的json反序列化方法
func (this *Gmt) UnmarshalJSON(data []byte) error {
	decoded := new(struct {
		GmtCreated  *JsonTime `json:"gmtCreated,omitempty"`
		GmtModified *JsonTime `json:"gmtModified,omitempty"`
	})
	err := json.Unmarshal(data, decoded)
	if err == nil {
		//this.GmtCreated = &decoded.GmtCreated

	}
	return err
}

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
