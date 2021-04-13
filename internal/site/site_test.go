package site

import (
	"fmt"
	"testing"

	"github.com/shoppehub/shoppe/datasource"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = datasource.Db
}

func TestCreate(t *testing.T) {

	db.AutoMigrate(&Site{})

	db.Create(&Site{Name: "D44", UserId: 100})

	var site2 Site

	db.Where("name = ?", "D44").Order("created_at desc").Find(&site2) // 查找 code 字段值为 D42 的记录

	fmt.Println(site2)
}
