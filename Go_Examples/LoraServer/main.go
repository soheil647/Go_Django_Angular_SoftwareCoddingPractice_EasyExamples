package main

import (
	"LoraServer/HandleRequests"
	processor "LoraServer/internal/request"
	"github.com/ghiac/go-commons/signal"
)


func main() {
	var modulesID []HandleRequests.ModuleID
	db := HandleRequests.SqlMigration("root:@/loraserver?charset=utf8&parseTime=True&loc=Local")
	jwt := HandleRequests.LoginRequest()
	db.Find(&modulesID)


	requestP := processor.NewRequest(db, jwt)
	requestP.Start(len(modulesID))
	for _, item := range modulesID {
		requestP.Tell(item.ModuleID)
	}

	signal.Signal.Wait()
}

//Send Data Service
/*MarshaledData := HandleRequests.MarshalSendData (data, item.ModuleID, db)
Status := HandleRequests.RequestSendData(MarshaledData, jwt, item.ModuleID)
if Status != nil{
	log.Fatal("Request Send failed", Status.Error())
} else {
	log.Fatal("Request Send Done")
}*/
