package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//note kalau nanti dah online localhost diganti sama ip server database
func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/Asap_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("gagal koneksi database")
	}

	db.AutoMigrate(&Predict{})

	return db
}
