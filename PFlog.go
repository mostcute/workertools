package workertools

import (
	"log"
	"os"
)

var  Pfloger PrintFileloger


type PrintFileloger interface {

	Println(a ...interface{})
	Printf(format string, v ...interface{})
}
type PrintFilelog struct {
	Flowlog   *log.Logger // 记录所有日志
	Printlog   *log.Logger // 记录所有日志
}
func (pf PrintFilelog)Println(a ...interface{})  {
	pf.Printlog.Println(a)
	pf.Flowlog.Println(a)
}
func (pf PrintFilelog)Printf(format string, v ...interface{})  {
	pf.Printlog.Printf(format,v)
	pf.Flowlog.Printf(format,v)
}

func init(){

	file, err := os.OpenFile("./managerlog.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}


	var Flowlog   *log.Logger // 记录所有日志
	var Printlog   *log.Logger // 记录所有日志

	Flowlog = log.New(file,"flow:",log.Ldate|log.Ltime|log.Lshortfile)
	Printlog = log.New(os.Stdout,"flow:",log.Ldate|log.Ltime|log.Lshortfile)
	Pflog :=PrintFilelog{Flowlog,Printlog}

	Pfloger = Pflog
}