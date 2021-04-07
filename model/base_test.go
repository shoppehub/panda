package model

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestJsonToString(t *testing.T) {
	// base := new(BaseStringId)

	now := time.Now()

	base := &BaseStringId{
		Id: "123",
		Gmt: Gmt{
			GmtCreated: &now,
		},
	}

	result, _ := json.Marshal(base)

	var newResult BaseStringId

	json.Unmarshal(result, &newResult)

	log.Println(string(result), newResult.Id, newResult.GmtCreated.String())
}
