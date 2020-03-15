package main

import (
	"fmt"
	models "github.com/cheyuexian/go-excise/orm/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func find(s *gorm.DB) error{

	var t1 []*models.T1
	err := s.Find(&t1).Error
	if err != nil{
		return  err
	}

	for _,t := range t1{
		fmt.Println(t)
	}
	return nil
}

func insert(s *gorm.DB) error{

	 t1 := models.T1{

		 Name:   randSeq(10),
		 Num:    rand.Int()%100,
		 Status: 0,
	 }
	 err := s.Create(&t1).Error

	return err
}
func update(s *gorm.DB) error{
	var t1 models.T1
	err := s.First(&t1).Where("id",1).Error
	if err != nil{
		fmt.Println(err)
		return err
	}

	fmt.Println(t1)

	//db.Model(&user).Where("active = ?", true).Update("name", "hello")
	err = s.Model(&models.T1{}).Update(models.T1{Id:1}).Update("num",gorm.Expr("num-?",1)).Error
	//s.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	//db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected
	//DB.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
////// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;
//https://gorm.io/docs/update.html
	if err != nil{
		fmt.Println(err)
		return err
	}

	return nil

}

func main() {
	rand.Seed(time.Now().Unix())
	db, err := gorm.Open("mysql", "root:123456@tcp(192.168.177.129:3306)/d1?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(err)

	tx := db.Begin()
	update(tx)
	//if err := find(tx);err!=nil {
	//	tx.Rollback()
	//	return
	//}
	if err := insert(tx);err!=nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	fmt.Println("after data")
	if err := find(tx);err!=nil {
		fmt.Println(err)

		tx.Rollback()
		return
	}

	tx.Commit()
	defer db.Close()
}
