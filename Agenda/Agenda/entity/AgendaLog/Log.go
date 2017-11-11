package AgendaLog

import (
	"log"
	"os"
)

func OperateLog(prefix string, info string) {
	file, err := os.OpenFile("Operate.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatalln("fail to open Operate.log file!")
	}
	defer file.Close()
	logger := log.New(file, prefix, log.LstdFlags|log.Llongfile)
	logger.Println(info)
}
