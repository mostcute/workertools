package workertools

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"os"
	"strconv"
	"time"
)

var Ch chan ReportWork
var Host string
var Ip string
var dia websocket.Dialer
var Con *websocket.Conn
var Server string
var Isinit bool
var IsManagerSet bool = false


var Workerlist map[string] Worker
type WorkerMsg struct {
	Msgtype string
	Data []byte
}

type ReportWork struct {
	Ip string
	Time int64
	Data WorkInfo
}
type WorkInfo struct {
	Status string
	Type string
	Sectors map[string]string
}
type Worker struct{
	Ip string
	Hostname string
	Time int64
	//tasks []uint64   //0:Addpiece 1:precommit1 2: precommit2 3:commit1 4:commit2
	Workerstatus  WorkerStatus
}
type WorkerStatus struct {
	Status string `json:"status"`
	Type string `json:"type"`
	Phrase map[string][]string `json:"phrase"`

}
func Workinit(){
	

	if IsManagerSet == false{//假设命令绝对正确，只存在无和绝对正确（无限重连）两种情况

		Pfloger.Println("don't receive manager address")

	}else {

		Ch = make(chan ReportWork,10000)
		var err error
		Host,err =os.Hostname()
		if err != nil{
			Pfloger.Println("Hostname:", err)
		}
		Ip = Get_hostip()
		Print_hostipall()
		if Isinit{
			return
		}
		Reconnect()
	}
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
		Pfloger.Println("dial:", err)
		return err
	}else{
		Pfloger.Println("reconnect success")
	}
	return nil
}
func CheckCh(){
	for{
		if len(Ch)>6000{
			Pfloger.Println("worker report Ch panic wraning")
			//TODO:save the ch in the location

		}

	}
}
func Processoutcoming(){
	if 	IsManagerSet == true{


		//tickTimer := time.NewTicker(7 * time.Second)
		go CheckCh()
		go func() {//目前不读取信息单纯维持心跳在线
			for{
				var Msg WorkerMsg
				err := Con.ReadJSON(Msg)
				if err != nil {
					Pfloger.Println("processIncomingMessage:", err)
					Reconnect()
				}
			}

		}()
		for {
			select {
			case  Msg :=<-Ch:
				js,err := json.Marshal(Msg)
				if err != nil {
					Pfloger.Println("processOutcomingMessage:", err)
				}
				var mmsg WorkerMsg
				mmsg.Msgtype = "Reportwork"
				mmsg.Data = js

				err = Con.WriteJSON(mmsg)
				if err != nil {
					Pfloger.Println("processOutcomingMessage:", err)
					Reconnect()
				}
				//case <- tickTimer.C://可重连
				//	var Msg msg.WorkerMsg
				//	err := Con.WriteJSON(Msg)
				//	if err != nil {
				//		Pfloger.Println("write:", err)
				//		Reconnect()
				//	}
			}
		}
	}


}
var ReportInfo WorkInfo

func Roport(phrase string, sid int64){
	if 	IsManagerSet == true {

		var Msg ReportWork
		Msg.Ip = Ip
		Msg.Time = time.Now().Unix()
		//ReportInfo.Type = "all"
		var sectorInfo = make(map[string]string)
		sectorInfo[phrase]=strconv.FormatInt(sid,10)
		ReportInfo.Sectors = sectorInfo
		Msg.Data = ReportInfo
		Ch <- Msg
	}

	//ret,_:= json.Marshal(Msg)
	//
	//fmt.Println(string(ret))
}


