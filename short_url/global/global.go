package global

import "log"

func InitGlobal() {
	log.Println("init MySQL")
	initMySQL()
}
