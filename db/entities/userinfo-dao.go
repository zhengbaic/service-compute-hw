package entities

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	db *gorm.DB
}

var mydb *DB

func GetDao() *DB {
	if mydb != nil {
		return mydb
	}
	db, err := gorm.Open("mysql", "root:root@tcp(192.168.99.100:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DataBase Connected")
		if !db.HasTable("User") {
			db.CreateTable(&User{})
		} else {
			db.DropTable(&User{})
			db.CreateTable(&User{})
		}
		mydb = &DB{db}
		return mydb
	}
}

func (userDB *DB) SaveUser(u *User) error {
	if err := userDB.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (userDB *DB) GetAUser(uid uint64) (*User, error) {
	var user User
	if err := userDB.db.Find(&user, uid).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (userDB *DB) GetAllUsers() ([]*User, error) {
	users := make([]*User, 0)
	if err := userDB.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
