package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/ErfanMomeniii/urlshortener-go/model"
	"errors"
)

func ConnectToDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{});
	if err!=nil {
		panic(err);
	}
	
	return db;
}

func HandleMigration(db *gorm.DB){
	if err:=db.AutoMigrate(&model.URL{});err!=nil{
		panic(err);
	}

	if err:=db.AutoMigrate(&model.User{});err!=nil{
		panic(err);
	}
}

func Create(model interface{},db *gorm.DB) error {
	if result:=db.Create(model);result.Error!=nil{
		return errors.New("model cant create");
	}
	
	return nil;
}

func Delete(model interface{},id int64,db *gorm.DB) error {
	if result:=db.Delete(model,id);result.Error!=nil{
		return errors.New("model cant delete");
	}
	
	return nil;
}