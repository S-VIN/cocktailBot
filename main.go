package main

import (
	"log"
	"os"
)

var logInf *log.Logger
var logErr *log.Logger

func main() {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	logInf = log.New(f, "INF\t", log.Ldate|log.Ltime)
	logErr = log.New(f, "ERR\t", log.Lshortfile)
	
	err = telegram.CreateBot()
	if err != nil {
		logErr.Panic(err)
	}
	
	err = telegram.CheckUpdates()
	if err != nil {
		logErr.Println(err)
	}
}
