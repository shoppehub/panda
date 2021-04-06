package model

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestJsonToString(t *testing.T) {
	// base := new(BaseStringId)

	ts, _ := time.Parse(JsonTimeLayout, "2021-04-06 14:18:09")
	fmt.Println(ts, "demo")

	customTime := JsonTime(time.Now())

	base := &BaseStringId{
		Id: "123",
		Gmt: Gmt{
			GmtCreated: customTime,
		},
	}

	ti := time.Time(customTime)

	fmt.Println(ti, 1)

	result, _ := json.Marshal(base)

	var newResult BaseStringId

	json.Unmarshal(result, &newResult)

	log.Println(string(result), newResult.Id, time.Time(newResult.GmtCreated))
}
