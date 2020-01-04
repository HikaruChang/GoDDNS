package util

import (
	"log"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		fileName, _ := os.OpenFile("GoDDNS.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.SetOutput(fileName)
		log.Println(err)
	}
}

func CheckErrCustom(err string) {
	if err != "" {
		fileName, _ := os.OpenFile("GoDDNS.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.SetOutput(fileName)
		log.Println(err)
	}
}
