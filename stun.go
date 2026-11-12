package main

import (
	"fmt"
	"log"

	"github.com/pion/stun/v3"
)

func stunlookup()(string,int,error){
	var pubip string 
	var pubport int 
	stunconn,err:=stun.Dial("udp4","stun.l.google.com:19302")
	
	if err!=nil{
		return "",0,err 
	}
	defer stunconn.Close()

	message:= stun.MustBuild(
		stun.TransactionID,
		stun.BindingRequest,
	)

	err=stunconn.Do(message,func(res stun.Event){
		if res.Error!=nil{
			log.Fatal(res.Error)
		}
		var mappedaddr stun.XORMappedAddress
		if err:= mappedaddr.GetFrom(res.Message); err!=nil{
			log.Fatal(err)
		}
		pubip =mappedaddr.IP.String()
		pubport = mappedaddr.Port
		fmt.Println("public ip: ",pubip)
		fmt.Println("public port: ",pubport)
	})
	if err!=nil{
		return "",0,err
	}
	return pubip , pubport , nil 	
}