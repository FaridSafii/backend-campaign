package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Pembuatan koneksi
	dsn := "root:@tcp(127.0.0.1:3306)/campaign_backend?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//Menampilkan kesalahan
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection to database success")
}
