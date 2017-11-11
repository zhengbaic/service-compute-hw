package User

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	UserName  string
	UserPass  string
	UserPhone string
	UserEmail string
}

func checkForRegister(body User, allUser []User) error {
	for _, value := range allUser {
		if value.UserName == body.UserName {
			return errors.New("This UserName has been used")
		} else if value.UserEmail == body.UserEmail {
			return errors.New("This Emial has been used")
		} else if value.UserPhone == body.UserPhone {
			return errors.New("This UserPhone has been used")
		}
	}
	return nil
}

func returnAllUser() []User {
	var allUser []User
	stream, _ := ioutil.ReadFile("User.json")
	json.Unmarshal(stream, &allUser)
	return allUser
}

func IsUser(userName string) bool {
	allUser := returnAllUser()
	for _, value := range allUser {
		if value.UserName == userName {
			return true
		}
	}
	return false
}

//增加注册用户
func UserRegitser(body User) error {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) != "" {
		return errors.New("You should louout the current account")
	}
	allUser := returnAllUser()
	err := checkForRegister(body, allUser)
	if err != nil {
		return err
	}
	allUser = append(allUser, body)
	result, _ := json.Marshal(allUser)
	file, _ := os.OpenFile("User.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file.Write(result)
	return nil
}

//登录
func UserLogin(userName, password string) error {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) != "" {
		return errors.New("Please logout the current account first")
	}
	allUser := returnAllUser()
	for index, value := range allUser {
		if value.UserName == userName && value.UserPass == password {
			file, _ := os.OpenFile("curUser.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
			file.WriteString(value.UserName)
			return nil
		} else if value.UserName == userName && value.UserPass != password {
			return errors.New("userName or password is wrong")
		} else if index == len(allUser)-1 {
			return errors.New("You has been not registered a account")
		}
	}
	return nil
}

//登出
func UserLogout() error {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) == "" {
		return errors.New("You has been logout")
	}
	file, _ := os.OpenFile("curUser.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file.WriteString("")
	return nil
}

//检测登录状态
func UserState() string {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) == "" {
		fmt.Println("You haven't logged in yet, Please login at first")
		return ""
	} else {
		return string(stream)
	}
}

//已登录用户删除用户
func UserDelete() error {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) == "" {
		return errors.New("Please login first")
	}
	allUser := returnAllUser()
	for index, value := range allUser {
		if value.UserName == string(stream) {
			allUser = append(allUser[:index], allUser[index+1:]...)
			break
		}
	}
	result, _ := json.Marshal(allUser)
	file, _ := os.OpenFile("User.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file.Write(result)
	file1, _ := os.OpenFile("curUser.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file1.WriteString("")
	return nil
}

func QueryUserByUserName() {
	allUser := returnAllUser()
	fmt.Println("The search result can be show as followed:")
	fmt.Println("UserName      " + "UserEmail     " + "UserPhone     ")
	for _, value := range allUser {
		fmt.Println(value.UserName + "   " + value.UserEmail + "   " + value.UserPhone)
	}
}
