package logger

import (
	"DouBanUpdater/config"
	"log"
)

func Success(str string, err error) {
	log.Println("[SUCCESS] " + str)
	if config.Config.Debug && err != nil {
		log.Println("[SUCCESS] Error: " + err.Error())
	}
}

func Info(str string, err error) {
	log.Println("[INFO] " + str)
	if config.Config.Debug && err != nil {
		log.Println("[INFO] Error: " + err.Error())
	}
}

func Warning(str string, err error) {
	log.Println("[WARN] " + str)
	if config.Config.Debug && err != nil {
		log.Println("[WARN] Error: " + err.Error())
	}
}

func Error(str string, err error) {
	log.Println("[ERROR] " + str)
	if config.Config.Debug && err != nil {
		log.Println("[ERROR] Error: " + err.Error())
	}
}

func Panic(str string, err error) {
	log.Println("[PANIC] " + str)
	if config.Config.Debug && err != nil {
		log.Println("[PANIC] Error: " + err.Error())
	}
}
