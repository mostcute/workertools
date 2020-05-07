package workertools

import (
	"encoding/json"
	"github.com/filecoin-project/yz-watch-manager/msg"
	"github.com/filecoin-project/yz-watch-manager/tools"
	"strconv"
	"time"
	"github.com/gorilla/websocket"
	"os"
)

var Ch chan msg.ReportWork
var Host string
var Ip string
var dia websocket.Dialer
var Con *websocket.Conn
var Server string
var Isinit bool
func Workinit(){
	Ch = make(chan msg.ReportWork,10000)
	var err error
	Host,err =os.Hostname()
	if err != nil{
		tools.Plog.Println("Hostname:", err)
	}
	Ip = tools.Get_hostip()
	tools.Print_hostipall()
	if Isinit{
		return
	}
	Reconnect()

	//Con, _, err = dia.Dial("ws://192.168.66.105:5678/worker", nil)
	//if err != nil {
	//	tools.Flog.Println("dial:", err)
	//	tools.Plog.Println("dial:", err)
	//}
}
func Reconnect(){
	var err error
	for{
		err = Connect()
		if err == nil{
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func Connect() error{
	var err error
	Con, _, err = dia.Dial("ws://"+Server+"/worker", nil)
	if err != nil {
		tools.Flog.Println("dial:", err)
		tools.Plog.Println("dial:", err)
		return err
	}else{
		tools.Flog.Println("reconnect success")
		tools.Plog.Println("reconnect success")
	}
	return nil
}
func CheckCh(){
	for{
		if len(Ch)>6000{
			tools.Flog.Println("worker report Ch panic wraning")
			tools.Plog.Println("worker report Ch panic wraning")
			//TODO:ȫ����������ػ���

		}

	}
}
func Processoutcoming(){

	//tickTimer := time.NewTicker(7 * time.Second)
	go func() {
		for{
			var Msg msg.WorkerMsg
			err := Con.ReadJSON(Msg)
			if err != nil {
				tools.Plog.Println("processIncomingMessage:", err)
				Reconnect()
			}
		}

	}()
	for {
		select {
		case  Msg :=<-Ch:
			js,err := json.Marshal(Msg)
			if err != nil {
				tools.Plog.Println("processOutcomingMessage:", err)
			}
			var mmsg msg.WorkerMsg
			mmsg.Msgtype = "Reportwork"
			mmsg.Data = js

			err = Con.WriteJSON(mmsg)
			if err != nil {
				tools.Plog.Println("processOutcomingMessage:", err)
				Reconnect()
			}
		//case <- tickTimer.C://可重连
		//	var Msg msg.WorkerMsg
		//	err := Con.WriteJSON(Msg)
		//	if err != nil {
		//		tools.Plog.Println("write:", err)
		//		Reconnect()
		//	}
		}
	}
}
var ReportInfo msg.WorkInfo

func Roport(phrase string, sid int64){
	var Msg msg.ReportWork
	Msg.Ip = Ip
	Msg.Time = time.Now().Unix()
	//ReportInfo.Type = "all"
	var sectorInfo = make(map[string]string)

	sectorInfo[phrase]=strconv.FormatInt(sid,10)
	
	ReportInfo.Sectors = sectorInfo
	Msg.Data = ReportInfo
	Ch <- Msg
	//ret,_:= json.Marshal(Msg)
	//
	//fmt.Println(string(ret))
}



func SendMsg(){

}