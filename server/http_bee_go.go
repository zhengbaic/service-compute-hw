package main

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

type MainC struct {
	beego.Controller
}

type FileC struct {
	beego.Controller
}

func (this *MainC) Get() {
	this.Ctx.WriteString("hello")
}

func (this *FileC) Get() {
	var age string
	if err := this.Ctx.Input.Bind(&age, "age"); err != nil {
		log.Println("no parameter")
	} else {
		log.Println("get [age] = ", age)
	}
	this.Data["json"] = map[string]interface{}{"age": age}
	this.ServeJSON()
	return
}

func main() {
	beego.BConfig.Listen.HTTPPort = 3001
	fmt.Println("hello beego")
	beego.Router("/", &MainC{})
	beego.Router("/temp", &FileC{})
	beego.Run()
}
